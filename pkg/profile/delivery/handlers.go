package profdel

import (
	"banana/pkg/domain"

	"github.com/gorilla/mux"
)

type profHandler struct {
	profUsecase domain.ProfUsecase
}

func SetProfHandlers(router *mux.Router, au domain.ProfUsecase) {
	handler := &profHandler{
		au,
	}

	router.HandleFunc(urlGetProfile, handler.GetProfile).Methods("GET", "OPTIONS")
}
