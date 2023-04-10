package reacdel

import (
	"banana/pkg/domain"
	"io/ioutil"

	"encoding/json"
	"net/http"
)

// /reactions/update
func (h *reacHandler) ReactionsUpd(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var emo domain.ReactionsApi
	err = json.Unmarshal(b, &emo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ReacUsecase.ReactionsUpd(emo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /question/ask
func (h *reacHandler) QuestionAsk(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.QuestionApi
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ReacUsecase.QuestionAsk(q.Hash, q.Option)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /question/like
func (h *reacHandler) QuestionLike(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.QuestionLikeApi
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ReacUsecase.QuestionLike(q.Hash, q.Idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
