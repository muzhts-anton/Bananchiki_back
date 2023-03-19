package presdel

import (
	"banana/pkg/domain"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

type getPresHttp struct {
	UserId int `json:"userId"`
	PresId int `json:"presId"`
}

func (h *presHandler) getPres(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var p getPresHttp
	err = json.Unmarshal(b, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// err = h.PresUsecase.GetPres(p) //
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (h *presHandler) createPres(w http.ResponseWriter, r *http.Request) {
	tmp := domain.Presentation{
		Url: "/tmp/pres",
		CreatorId: 35152,
	}
	h.PresUsecase.UploadPres(&tmp)
	w.WriteHeader(http.StatusOK)
}
