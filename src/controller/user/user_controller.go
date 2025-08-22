package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	manager_error "github.com/LuizFreitas225/user-manager-go/src/configuration/rest_erro"
	httpConstant "github.com/LuizFreitas225/user-manager-go/src/controller/constant"
	"github.com/LuizFreitas225/user-manager-go/src/controller/user/data"
	"github.com/LuizFreitas225/user-manager-go/src/repository/user"
	"github.com/gorilla/mux"
)

type UserController struct {
	Repository user.Repository
}

func (*UserController) Create(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	input := data.InputUserOfFindById{}

	n, err := strconv.Atoi(idStr)
	if err != nil {
		reposeError := manager_error.NewBadRequestError("Error converting id to integer.")
		uc.writeError(&w, reposeError, reposeError.Code)
		return
	}
	input.ID = n

	result, err := uc.Repository.FindById(input)

	if err != nil {
		uc.writeError(&w, err, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
		json, _ := json.Marshal(result)
		w.Write(json)
	}
}

func (uc *UserController) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	input := data.InputUserOfSearch{}

	if query.Has("searchTerm") {
		input.SearchTerm = query.Get("searchTerm")
	}

	result, err := uc.Repository.Search(input)

	if err != nil {
		uc.writeError(&w, err, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
		json, _ := json.Marshal(result)
		w.Write(json)
	}
}

func (*UserController) Update(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	input := data.InputUserOfDelete{}

	n, err := strconv.Atoi(idStr)
	if err != nil {
		reposeError := manager_error.NewBadRequestError("Error converting id to integer.")
		uc.writeError(&w, reposeError, reposeError.Code)
		return
	}
	input.ID = n

	err = uc.Repository.Delete(input)

	if err != nil {
		if restErr, ok := err.(*manager_error.RestError); ok {
			uc.writeError(&w, restErr, restErr.Code)
		} else {
			uc.writeError(&w, err, http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
	}
}

func (*UserController) writeError(w *http.ResponseWriter, err error, status int) {
	writer := *w
	writer.WriteHeader(status)
	writer.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
	json, _ := json.Marshal(err)
	writer.Write(json)
}
