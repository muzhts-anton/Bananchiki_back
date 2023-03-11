package main

import (
	"github.com/gorilla/mux"

	"banana/pkg/show/delivery"
	"banana/pkg/show/usecase"
	"banana/pkg/show/repository"

	"fmt"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	quizRep := quizrep.InitQuizRep()
	quizUsc := quizusc.InitAnnUsc(quizRep)
	quizdel.SetQuizHandlers(api, quizUsc)

	fmt.Println("start")
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}