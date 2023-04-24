package quizdel

import (
	"banana/pkg/domain"

	"github.com/gorilla/mux"
)

type quizHandler struct {
	QuizUsecase domain.QuizUsecase
}

func SetQuizHandlers(router *mux.Router, uc domain.QuizUsecase) {
	handler := &quizHandler{
		QuizUsecase: uc,
	}

	router.HandleFunc(urlCreateQuiz, handler.createQuiz).Methods("POST", "OPTIONS")
	router.HandleFunc(urlDeleteQuiz, handler.deleteQuiz).Methods("POST", "OPTIONS")
	router.HandleFunc(urlUpdateQuiz, handler.updateQuiz).Methods("PUT", "OPTIONS")

	router.HandleFunc(urlCreateQuizVote, handler.createQuizVote).Methods("POST", "OPTIONS")
	router.HandleFunc(urlDeleteQuizVote, handler.deleteQuizVote).Methods("POST", "OPTIONS")
	router.HandleFunc(urlUpdateQuizVote, handler.updateQuizVote).Methods("PUT", "OPTIONS")
	router.HandleFunc(urlPollQuizVote, handler.pollQuizVote).Methods("PUT", "OPTIONS")

	router.HandleFunc(urlCompetitionStart, handler.competitionStart).Methods("POST", "OPTIONS")
	router.HandleFunc(urlCompetitionStop, handler.competitionStop).Methods("POST", "OPTIONS")
	router.HandleFunc(urlCompetitionVoterRegister, handler.competitionVoterRegister).Methods("POST", "OPTIONS")
	router.HandleFunc(urlCompetitionResult, handler.GetCompetitionResult).Methods("GET", "OPTIONS")
}
