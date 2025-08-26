package user

import "github.com/LuizFreitas225/user-manager-go/src/controller/user/data"

type Repository interface {
	Search(input data.InputUserOfSearch) ([]data.OutputUserOfSearch, error)
	FindById(input data.InputUserOfFindById) (data.OutputUserOfFindById, error)
	Delete(input data.InputUserOfDelete) error
	Create(input data.InputUserOfCreate) (data.OutputUserOfCreate, error)
	Update(input data.InputUserOfUpdate) (data.OutputUserOfUpdate, error)
}
