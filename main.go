package main

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"banana/pkg/database"
	presgrpc "banana/pkg/presentation/delivery/grpc"

	"banana/pkg/quiz/delivery"
	"banana/pkg/quiz/repository"
	"banana/pkg/quiz/usecase"

	"banana/pkg/presentation/delivery"
	"banana/pkg/presentation/repository"
	"banana/pkg/presentation/usecase"

	"fmt"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

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
