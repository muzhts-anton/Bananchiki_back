package middlewares

import (
	"banana/pkg/utils/log"

	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.InfoWithoutCaller(fmt.Sprintf("[%s] %s, %s %s",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start)))
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Warn(fmt.Sprintf("Recovered from panic with err: %s on %s", err, r.RequestURI))
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://kindaslides.ru/") // url to deployed front
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, Location")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Max-Age", "600")
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

var Csrf = csrf.Protect(
	[]byte("32-byte-long-auth-key"),
	csrf.Path("/"),
	csrf.Secure(false),
)
