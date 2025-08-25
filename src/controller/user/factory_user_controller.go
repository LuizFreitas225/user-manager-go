package user

import (
	"github.com/LuizFreitas225/user-manager-go/src/repository/user"
	"github.com/go-playground/validator"
)

func CreateUserController() *UserController {
	return &UserController{
		Repository: user.CreateUserRepository(),
		Validate:   *validator.New(),
	}
}
