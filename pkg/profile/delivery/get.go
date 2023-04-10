package profdel

import (
	"banana/pkg/utils/sessions"

	"encoding/json"
	"net/http"
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
