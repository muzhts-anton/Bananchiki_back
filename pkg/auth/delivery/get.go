package authdel

import (
	"banana/pkg/domain"
	"banana/pkg/utils/sessions"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

// /user/register
func (handler *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userForm := new(domain.User)
	err = json.Unmarshal(b, userForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	us, err := handler.AuthUsecase.Register(domain.User{
		Id:             userForm.Id,
		Username:       userForm.Username,
		Password:       userForm.Password,
		Email:          userForm.Email,
		Imgsrc:         userForm.Imgsrc,
		RepeatPassword: userForm.RepeatPassword,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sessions.StartSession(w, r, us.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

// /user/login
func (handler *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userForm := new(domain.UserBasic)
	err = json.Unmarshal(b, userForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	us, err := handler.AuthUsecase.Login(domain.UserBasic{
		Email:    userForm.Email,
		Password: userForm.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = sessions.CheckSession(r); err != domain.ErrUserNotLoggedIn {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = sessions.StartSession(w, r, us.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// /user/logout
func (handler *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if err := sessions.FinishSession(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /user/session
func (handler *authHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	_, err := sessions.CheckSession(r)
	if err == domain.ErrUserNotLoggedIn {
		out, err := json.Marshal(domain.SessionInfo{Status: false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(domain.SessionInfo{Status: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
