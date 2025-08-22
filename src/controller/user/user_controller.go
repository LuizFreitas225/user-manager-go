package user

import (
	"encoding/json"
	"net/http"

	httpConstant "github.com/LuizFreitas225/user-manager-go/src/controller/constant"
	"github.com/LuizFreitas225/user-manager-go/src/controller/user/data"
	"github.com/LuizFreitas225/user-manager-go/src/repository/user"
)

type UserController struct {
	Repository user.Repository
}

func (*UserController) Create(w http.ResponseWriter, r *http.Request) {

}

func (*UserController) FindById(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	input := data.InputUserOfSearch{}

	if query.Has("searchTerm") {
		input.SearchTerm = query.Get("searchTerm")
	}

	result, err := uc.Repository.Search(input)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
		json, _ := json.Marshal(err)
		w.Write(json)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
		json, _ := json.Marshal(result)
		w.Write(json)
	}
}

func (*UserController) Update(w http.ResponseWriter, r *http.Request) {

}

func (*UserController) Delete(w http.ResponseWriter, r *http.Request) {

}
