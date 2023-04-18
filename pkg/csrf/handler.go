package csrf

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func SetHandler(api *mux.Router) {
	api.HandleFunc("/csrf", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		w.WriteHeader(http.StatusNoContent)
	}).Methods("GET", "OPTIONS")
}
