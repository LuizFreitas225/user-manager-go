package user

import "net/http"

type Controller interface {
	Create(http.ResponseWriter, *http.Request)
	Find(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
