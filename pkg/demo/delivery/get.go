package demodel

import (
	"banana/pkg/domain"
	"banana/pkg/utils/hash"

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