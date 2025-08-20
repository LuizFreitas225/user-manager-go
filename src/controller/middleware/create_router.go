package middleware

import (
	"github.com/LuizFreitas225/user-manager-go/src/controller/login"
	"github.com/LuizFreitas225/user-manager-go/src/controller/user"
	"github.com/LuizFreitas225/user-manager-go/src/system/singleton"
)

func CreateRouterManager() ManagerRouter {
	return ManagerRouter{
		Router:          singleton.GetInstance().Router,
		UserController:  &user.UserController{},
		LoginController: &login.LoginController{},
	}
}
