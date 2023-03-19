package presdel

import (
	"banana/pkg/domain"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type presApiRequest struct {
	CreatorId uint64 `json:"creatorId"`
}

// /presentation/{id}
func (h *presHandler) getPres(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	presId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var p presApiRequest
	err = json.Unmarshal(b, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pres, err := h.PresUsecase.GetPres(p.CreatorId, presId)
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
	tmp := domain.Presentation{
		Url:       "/tmp/pres",
		CreatorId: 35152,
	}
	h.PresUsecase.CreatePres(&tmp)
	w.WriteHeader(http.StatusOK)
}
