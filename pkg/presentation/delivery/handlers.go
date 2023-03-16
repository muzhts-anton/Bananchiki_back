package presdel

import (
	"banana/pkg/presentation/usecase"

	"github.com/gorilla/mux"
)

type PresHandler struct {
	PresUsecase *presusc.PresUsecase
}

func SetPresHandlers(router *mux.Router, uc *presusc.PresUsecase) {
	handler := &PresHandler{
		PresUsecase: uc,
	}

	router.HandleFunc(getPres, handler.getPres).Methods("GET", "OPTIONS")
	router.HandleFunc(uploadPres, handler.uploadPres).Methods("POST", "OPTIONS")
}