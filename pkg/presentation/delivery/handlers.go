package presdel

import (
	"banana/pkg/domain"

	"github.com/gorilla/mux"
)

type PresHandler struct {
	PresUsecase domain.PresUsecase
}

func SetQuizHandlers(router *mux.Router, uc domain.PresUsecase) {
	handler := &PresHandler{
		PresUsecase: uc,
	}

	router.HandleFunc(getPres, handler.getPres).Methods("GET", "OPTIONS")
	router.HandleFunc(uploadPres, handler.uploadPres).Methods("POST", "OPTIONS")
}