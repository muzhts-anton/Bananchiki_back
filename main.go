package main

import (
	"banana/pkg/database"
	presgrpc "banana/pkg/presentation/delivery/grpc"
	"banana/pkg/utils/log"

	"banana/pkg/quiz/delivery"
	"banana/pkg/quiz/repository"
	"banana/pkg/quiz/usecase"

	"banana/pkg/presentation/delivery"
	"banana/pkg/presentation/repository"
	"banana/pkg/presentation/usecase"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"fmt"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(Logger)
	api.Use(PanicRecovery)

	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	quizRep := quizrep.InitQuizRep(db)
	quizUsc := quizusc.InitQuizUsc(quizRep)
	quizdel.SetQuizHandlers(api, quizUsc)

	conn, _ := grpc.Dial(":50051", grpc.WithInsecure())
	c := presgrpc.NewParsingClient(conn)

	presRepo := presrep.InitPresRep(db)
	presUsc := presusc.InitPresUscase(c, presRepo)
	presdel.SetPresHandlers(api, presUsc)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

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
		w.Header().Set("Access-Control-Allow-Origin", "http://185.241.192.112/") // url to deployed front
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
