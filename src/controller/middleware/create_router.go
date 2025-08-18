package middleware

import (
	"github.com/LuizFreitas225/user-manager-go/src/controller/login"
	"github.com/LuizFreitas225/user-manager-go/src/controller/user"
	"github.com/gorilla/mux"
)

func CreateRouterManager() ManagerRouter {
	return ManagerRouter{
		Router:          mux.NewRouter(),
		UserController:  &user.UserController{},
		LoginController: &login.LoginController{},
	}
}
