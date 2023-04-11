package presdel

import (
	"banana/pkg/presentation/usecase"

	"github.com/gorilla/mux"
)

type presHandler struct {
	PresUsecase *presusc.PresUsecase
}

func SetPresHandlers(router *mux.Router, uc *presusc.PresUsecase) {
	handler := &presHandler{
		PresUsecase: uc,
	}

	router.HandleFunc(urlGetPres, handler.getPres).Methods("GET", "OPTIONS")
	router.HandleFunc(urlCreatePres, handler.createPres).Methods("POST", "OPTIONS")
	router.HandleFunc(urlPresNameChange, handler.changePresName).Methods("PUT", "OPTIONS")
}
