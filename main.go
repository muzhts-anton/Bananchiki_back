package main

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	tmp "banana/pkg/presentation/delivery/grpc"

	"banana/pkg/quiz/delivery"
	"banana/pkg/quiz/repository"
	"banana/pkg/quiz/usecase"

	"banana/pkg/presentation/delivery"
	"banana/pkg/presentation/usecase"

	"fmt"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	quizRep := quizrep.InitQuizRep()
	quizUsc := quizusc.InitAnnUsc(quizRep)
	quizdel.SetQuizHandlers(api, quizUsc)

	conn, _ := grpc.Dial(":50051", grpc.WithInsecure())
	c := tmp.NewParsingClient(conn)

	t1 := presusc.InitPresUscase(c)
	presdel.SetPresHandlers(api, t1)

	fmt.Println("start")
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}