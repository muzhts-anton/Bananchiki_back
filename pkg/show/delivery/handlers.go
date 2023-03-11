package quizdel

import (
	"banana/pkg/domain"

	"github.com/gorilla/mux"
)

type QuizHandler struct {
	QuizUsecase domain.QuizUsecase
}

func SetQuizHandlers(router *mux.Router, uc domain.QuizUsecase) {
	handler := &QuizHandler{
		QuizUsecase: uc,
	}

	router.HandleFunc(createQuiz, handler.createQuiz).Methods("POST", "OPTIONS")
	router.HandleFunc(voteQuiz, handler.voteQuiz).Methods("POST", "OPTIONS")
	router.HandleFunc(showQuiz, handler.showQuiz).Methods("GET", "OPTIONS")
}