package repository

import "github.com/LuizFreitas225/user-manager-go/src/system/singleton"

func CreateUserRepository() UserRepository {
	return UserRepository{
		Db: singleton.GetInstance().Db,
	}
}
