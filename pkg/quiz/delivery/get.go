package quizdel

import (
	"banana/pkg/domain"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

// /quiz/create
func (h *quizHandler) createQuiz(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q := domain.QuizHTTP{}
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	qid, err := h.QuizUsecase.CreateQuiz(q.Quiz, q.CreatorId, q.PresId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(struct {
		Id uint64 `json:"quizId"`
	}{Id: qid},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// /quiz/update
func (h *quizHandler) updateQuiz(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.QuizHTTP
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.QuizUsecase.UpdateQuiz(q.Quiz, q.CreatorId, q.PresId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// quiz/delete
func (h *quizHandler) deleteQuiz(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.QuizHTTP
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.QuizUsecase.DeleteQuiz(q.Quiz.Id, q.CreatorId, q.PresId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// quiz/vote/delete
func (h *quizHandler) deleteQuizVote(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.VoteHTTP
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.QuizUsecase.DeleteQuizVote(q.Vote.Idx, q.QuizId, q.CreatorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /quiz/vote/update
func (h *quizHandler) updateQuizVote(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.QuizHTTP
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	err = h.QuizUsecase.UpdateQuizVote(q.Quiz.Votes, q.Quiz.Id, q.CreatorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /quiz/vote/create
func (h *quizHandler) createQuizVote(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.VoteHTTP
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.QuizUsecase.CreateQuizVote(q.Vote, q.QuizId, q.CreatorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /quiz/vote/poll
func (h *quizHandler) pollQuizVote(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.VoteHTTP
	err = json.Unmarshal(b, &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.QuizUsecase.PollQuizVote(q.Vote.Idx, q.QuizId, q.VoterName, q.VoterId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *quizHandler) competitionStart(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.CompetitionHttp
	err = json.Unmarshal(b, &q)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.QuizUsecase.CompetitionStart(q.QuizId, q.PresId)

	w.WriteHeader(http.StatusOK)
}

func (h* quizHandler) competitionStop(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.CompetitionHttp
	err = json.Unmarshal(b, &q)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.QuizUsecase.CompetitionStop(q.QuizId, q.PresId)

	w.WriteHeader(http.StatusOK)
}

func (h* quizHandler) competitionVoterRegister(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q domain.VoteRegisterHttp
	err = json.Unmarshal(b, &q)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vId, err := h.QuizUsecase.CompetitionVoterRegister(q.Name, q.Hash)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(struct {
		Id uint64 `json:"id"`
	}{Id: vId},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (h* quizHandler) GetCompetitionResult(w http.ResponseWriter, r *http.Request){
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

	newResults, err := h.QuizUsecase.GetCompetitionResult(presId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(struct {
		Top []domain.ResultItem `json:"top"`
	}{Top: newResults},
	)
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}