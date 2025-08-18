package middleware

import (
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

func (managerRouter *ManagerRouter) Start() {

	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"})
	// handler := httpgzip.NewHandler(managerRouter.Router, nil)
	// handler = handlers.CORS(originsOk, methodsOk)(handler)

	managerRouter.InitRoutes()

	go func() {
		err := http.ListenAndServe("localhost:8080", managerRouter.Router)

		if err != nil {
			println("Erro ao executar aplicação: ", err)
		}
	}()

}

func (managerRouter *ManagerRouter) InitRoutes() {
	managerRouter.initUserRoutes()
	managerRouter.initLoginRoutes()
}

func (managerRouter *ManagerRouter) initUserRoutes() {
	userSubRouter := managerRouter.Router.PathPrefix("/user").Subrouter()

	userSubRouter.HandleFunc("/create/:userId", managerRouter.UserController.Find).Methods(http.MethodPost)
	userSubRouter.HandleFunc("/find/", managerRouter.UserController.Find).Methods(http.MethodGet)
	userSubRouter.HandleFunc("/update/:userId", managerRouter.UserController.Find).Methods(http.MethodPut)
	userSubRouter.HandleFunc("/delete/:userId", managerRouter.UserController.Find).Methods(http.MethodDelete)
}

func (managerRouter *ManagerRouter) initLoginRoutes() {
	userSubRouter := managerRouter.Router.PathPrefix("/login").Subrouter()

	userSubRouter.HandleFunc("", managerRouter.LoginController.Login).Methods(http.MethodPost)
}
