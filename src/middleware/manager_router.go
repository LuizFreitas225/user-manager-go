package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/LuizFreitas225/user-manager-go/src/controller/login"
	"github.com/LuizFreitas225/user-manager-go/src/controller/user"
	"github.com/gorilla/mux"
)

type ManagerRouter struct {
	Router          *mux.Router
	UserController  user.Controller
	LoginController login.Controller
}

func (managerRouter *ManagerRouter) Start(ctx context.Context) {

	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"})
	// handler := httpgzip.NewHandler(managerRouter.Router, nil)
	// handler = handlers.CORS(originsOk, methodsOk)(handler)

	managerRouter.InitRoutes()

	// Cria o servidor HTTP
	server := &http.Server{
		Addr:    ":8080",
		Handler: managerRouter.Router,
	}

	//Roda o servidor dentro de uma Goroutine
	go func() {
		log.Println("Servidor HTTP iniciado na porta 8080")
		err := server.ListenAndServe()

		if err != nil {
			println(err.Error())
		}
	}()

	<-ctx.Done()
	log.Println("Contexto cancelado. Desligando servidor...")

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Erro ao desligar servidor: %v", err)
	} else {
		log.Println("Servidor desligado com sucesso")
	}
}

func (managerRouter *ManagerRouter) InitRoutes() {
	managerRouter.initUserRoutes()
	managerRouter.initLoginRoutes()
}

func (managerRouter *ManagerRouter) initUserRoutes() {
	userSubRouter := managerRouter.Router.PathPrefix("/user").Subrouter()

	userSubRouter.HandleFunc("/create/", managerRouter.UserController.Find).Methods(http.MethodPost)
	userSubRouter.HandleFunc("/find/", managerRouter.UserController.Find).Methods(http.MethodGet)
	userSubRouter.HandleFunc("/update/:userId", managerRouter.UserController.Find).Methods(http.MethodPut)
	userSubRouter.HandleFunc("/delete/:userId", managerRouter.UserController.Find).Methods(http.MethodDelete)
}

func (managerRouter *ManagerRouter) initLoginRoutes() {
	userSubRouter := managerRouter.Router.PathPrefix("/login").Subrouter()

	userSubRouter.HandleFunc("", managerRouter.LoginController.Login).Methods(http.MethodPost)
}
