package user

import "github.com/LuizFreitas225/user-manager-go/src/repository/user"

func CreateUserController() *UserController {
	return &UserController{
		Repository: user.CreateUserRepository(),
	}
}
