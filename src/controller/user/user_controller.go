package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	manager_error "github.com/LuizFreitas225/user-manager-go/src/configuration/rest_erro"
	httpConstant "github.com/LuizFreitas225/user-manager-go/src/controller/constant"
	"github.com/LuizFreitas225/user-manager-go/src/controller/user/data"
	"github.com/LuizFreitas225/user-manager-go/src/repository/user"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type UserController struct {
	Repository user.Repository
	Validate   validator.Validate
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var input data.InputUserOfCreate

	// Decodifica o JSON do body dentro do input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := uc.Validate.Struct(input); err != nil {
		uc.writeError(&w, manager_error.NewBadRequestError(err.Error()), http.StatusBadRequest)
		return
	}

	createdUser, err := uc.Repository.Create(input)
	if err != nil {
		if restErr, isRestErr := err.(*manager_error.RestError); isRestErr {
			uc.writeError(&w, err, restErr.Code)
		} else {
			uc.writeError(&w, err, http.StatusInternalServerError)
		}
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
		json, _ := json.Marshal(createdUser)
		w.Write(json)
	}
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
		if restErr, isRestErr := err.(*manager_error.RestError); isRestErr {
			uc.writeError(&w, err, restErr.Code)
		} else {
			uc.writeError(&w, err, http.StatusInternalServerError)
		}
		return
	} else {
		uc.writeResultOK(w, result)
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
		if restErr, isRestErr := err.(*manager_error.RestError); isRestErr {
			uc.writeError(&w, err, restErr.Code)
		} else {
			uc.writeError(&w, err, http.StatusInternalServerError)
		}
		return
	} else {
		uc.writeResultOK(w, result)
	}
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var input data.InputUserOfUpdate

	// Decodifica o JSON do body dentro do input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := uc.Validate.Struct(input); err != nil {
		uc.writeError(&w, manager_error.NewBadRequestError(err.Error()), http.StatusBadRequest)
		return
	}

	updatedUser, err := uc.Repository.Update(input)
	if err != nil {
		if restErr, isRestErr := err.(*manager_error.RestError); isRestErr {
			uc.writeError(&w, err, restErr.Code)
		} else {
			uc.writeError(&w, err, http.StatusInternalServerError)
		}
		return
	} else {
		uc.writeResultOK(w, updatedUser)
	}
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

func (*UserController) writeResultOK(w http.ResponseWriter, result any) {
	w.Header().Add(httpConstant.HeaderContentType, httpConstant.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Erro to convert resutl in json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
