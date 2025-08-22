package user

import "github.com/LuizFreitas225/user-manager-go/src/controller/user/data"

type Repository interface {
	Search(input data.InputUserOfSearch) ([]data.OutputUserOfSearch, error)
}
