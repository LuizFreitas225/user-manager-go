package login

import "net/http"

type Controller interface {
	Login(http.ResponseWriter, *http.Request)
}
