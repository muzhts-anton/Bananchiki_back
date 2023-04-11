package profdel

import (
	"banana/pkg/utils/sessions"
	"banana/pkg/domain"
	"banana/pkg/utils/filesaver"

	"path/filepath"
	"encoding/json"
	"net/http"
	"banana/pkg/utils/log"
)

// /profile
func (h *profHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	usrId, err := sessions.CheckSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prof, err := h.profUsecase.GetProfile(usrId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(prof)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (h* profHandler) UpdateProfileAvatar(w http.ResponseWriter, r* http.Request){
	userId, err := sessions.CheckSession(r)
	if err == domain.ErrUserNotLoggedIn {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.ParseMultipartForm(3 * 1024 * 1024) // limit 10Mb
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	

	uploaded, header, err := r.FormFile("avatar")
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploaded.Close()

	filename, err := filesaver.UploadFile(uploaded, domain.AvatarFilePath, filepath.Ext(header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.profUsecase.UpdateProfileAvatar(filename, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := domain.AvatarFilePath + filename
	out, err := json.Marshal(struct {
		Path string `json:"path"`
	}{Path: path})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
