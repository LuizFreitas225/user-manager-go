package user

import "net/http"

type Controller interface {
	Create(http.ResponseWriter, *http.Request)
	FindById(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)	
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
