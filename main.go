package main

import (
	"banana/pkg/csrf"
	presgrpc "banana/pkg/presentation/delivery/grpc"
	"banana/pkg/utils/database"
	"banana/pkg/utils/log"
	"banana/pkg/utils/middlewares"

	"banana/pkg/profile/delivery"
	"banana/pkg/profile/repository"
	"banana/pkg/profile/usecase"

	"banana/pkg/quiz/delivery"
	"banana/pkg/quiz/repository"
	"banana/pkg/quiz/usecase"

	"banana/pkg/auth/delivery"
	"banana/pkg/auth/repository"
	"banana/pkg/auth/usecase"

	"banana/pkg/presentation/delivery"
	"banana/pkg/presentation/repository"
	"banana/pkg/presentation/usecase"

	"banana/pkg/demo/delivery"
	"banana/pkg/demo/repository"
	"banana/pkg/demo/usecase"

	"banana/pkg/reactions/delivery"
	"banana/pkg/reactions/repository"
	"banana/pkg/reactions/usecase"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"fmt"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(middlewares.Logger)
	api.Use(middlewares.Cors)
	api.Use(middlewares.PanicRecovery)
	api.Use(middlewares.Csrf)

	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	profRep := profrep.InitProfRep(db)
	profUsc := profusc.InitProfUsc(profRep)
	profdel.SetProfHandlers(api, profUsc)

	quizRep := quizrep.InitQuizRep(db)
	quizUsc := quizusc.InitQuizUsc(quizRep)
	quizdel.SetQuizHandlers(api, quizUsc)

	authRep := authrep.InitAuthRep(db)
	authUsc := authusc.InitAuthUsc(authRep)
	authdel.SetAuthHandlers(api, authUsc)

	demoRep := demorep.InitDemoRep(db)
	demoUsc := demousc.InitDemoUsc(demoRep)
	demodel.SetAuthHandlers(api, demoUsc)

	reacRep := reacrep.InitReacRep(db)
	reacUsc := reacusc.InitReacUsc(reacRep)
	reacdel.SetReacHandlers(api, reacUsc)

	conn, _ := grpc.Dial(":50051", grpc.WithInsecure())
	c := presgrpc.NewParsingClient(conn)

	presRepo := presrep.InitPresRep(db)
	presUsc := presusc.InitPresUscase(c, presRepo)
	presdel.SetPresHandlers(api, presUsc)

	csrf.SetHandler(api)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: router,
	}

	log.Info("connecting to :3000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
