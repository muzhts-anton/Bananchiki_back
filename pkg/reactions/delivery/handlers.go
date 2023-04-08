package reacdel

import (
	"banana/pkg/domain"

	"github.com/gorilla/mux"
)

type reacHandler struct {
	ReacUsecase domain.ReacUsecase
}

func SetReacHandlers(router *mux.Router, au domain.ReacUsecase) {
	handler := &reacHandler{
		au,
	}

	router.HandleFunc(urlReactionsUpd, handler.ReactionsUpd).Methods("PUT", "OPTIONS")
}
