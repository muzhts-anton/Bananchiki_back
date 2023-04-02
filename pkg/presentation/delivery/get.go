package presdel

import (
	"banana/pkg/domain"
	"banana/pkg/utils/filesaver"
	"banana/pkg/utils/log"
	"banana/pkg/utils/sessions"

	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

// /presentation/{id}
func (h *presHandler) getPres(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	presId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	usrId, err := sessions.CheckSession(r)
	if err == domain.ErrUserNotLoggedIn {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pres, err := h.PresUsecase.GetPres(usrId, presId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(struct {
		Pres domain.PresApiResponse `json:"pres"`
	}{Pres: pres})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// /presentation/create
func (h *presHandler) createPres(w http.ResponseWriter, r *http.Request) {
	usrId, err := sessions.CheckSession(r)
	if err == domain.ErrUserNotLoggedIn {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = r.ParseMultipartForm(10 * 1024 * 1024) // limit 10Mb
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploaded, header, err := r.FormFile("presentation")
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploaded.Close()

	filename, err := filesaver.UploadFile(uploaded, domain.PresentationFilePath, filepath.Ext(header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	presId, err := h.PresUsecase.CreatePres(filename, usrId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(struct {
		PresId uint64 `json:"presId"`
	}{PresId: presId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
