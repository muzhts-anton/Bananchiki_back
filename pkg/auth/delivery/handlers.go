package authdel

import (
	"banana/pkg/domain"

	"github.com/gorilla/mux"
)

type authHandler struct {
	AuthUsecase domain.AuthUsecase
}

func SetAuthHandlers(router *mux.Router, au domain.AuthUsecase) {
	handler := &authHandler{
		au,
	}

	router.HandleFunc(urlReg, handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc(urlLogin, handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc(urlLogout, handler.Logout).Methods("PUT", "OPTIONS")
	router.HandleFunc(urlGetSession, handler.GetSession).Methods("GET", "OPTIONS")
}
