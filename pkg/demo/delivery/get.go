package demodel

import (
	"banana/pkg/domain"
	"banana/pkg/utils/hash"
	"banana/pkg/utils/sessions"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// /presentation/view/join/{code}
func (handler *demoHandler) JoinDemo(w http.ResponseWriter, r *http.Request) {
	code, ok := mux.Vars(r)["code"]
	if !ok || len(code) != 4 {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(struct {
		Hash string `json:"hash"`
	}{Hash: hash.EncodeToHash(code)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// /presentation/view/{hash}
func (handler *demoHandler) ViewDemo(w http.ResponseWriter, r *http.Request) {
	urlHash, ok := mux.Vars(r)["hash"]
	if !ok {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	cs, err := handler.DemoUsecase.GetCurrentDemoSlide(urlHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// /presentation/{presId}/show/go/{idx}
func (handler *demoHandler) ShowDemoGo(w http.ResponseWriter, r *http.Request) {
	presIdStr, ok := mux.Vars(r)["presId"]
	if !ok {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}
	presId, err := strconv.ParseUint(presIdStr, 10, 64)
	if err != nil {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	idxStr, ok := mux.Vars(r)["idx"]
	if !ok {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}
	idx, err := strconv.ParseUint(idxStr, 10, 32)
	if err != nil {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	userId, err := sessions.CheckSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.DemoUsecase.ShowDemoGo(presId, userId, uint32(idx))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /presentation/{presId}/show/stop
func (handler *demoHandler) ShowDemoStop(w http.ResponseWriter, r *http.Request) {
	presIdStr, ok := mux.Vars(r)["presId"]
	if !ok {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}
	presId, err := strconv.ParseUint(presIdStr, 10, 64)
	if err != nil {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	userId, err := sessions.CheckSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.DemoUsecase.ShowDemoSop(presId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
