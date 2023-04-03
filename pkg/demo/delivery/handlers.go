package demodel

import (
	"banana/pkg/domain"

	"github.com/gorilla/mux"
)

type demoHandler struct {
	DemoUsecase domain.DemoUsecase
}

func SetAuthHandlers(router *mux.Router, au domain.DemoUsecase) {
	handler := &demoHandler{
		au,
	}

	router.HandleFunc(urlViewJoin, handler.JoinDemo).Methods("GET", "OPTIONS")
	router.HandleFunc(urlView, handler.ViewDemo).Methods("GET", "OPTIONS")
	router.HandleFunc(urlShowGo, handler.ShowDemoGo).Methods("PUT", "OPTIONS")
	router.HandleFunc(urlShowStop, handler.ShowDemoStop).Methods("PUT", "OPTIONS")
}
