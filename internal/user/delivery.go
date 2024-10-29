package user

import "net/http"

type Handlers interface {
	Health(response http.ResponseWriter, request *http.Request)
	Register(response http.ResponseWriter, request *http.Request)
	Login(response http.ResponseWriter, request *http.Request)
	Logout(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	UploadImage(response http.ResponseWriter, request *http.Request)
	GetUserByUsername(response http.ResponseWriter, request *http.Request)
	GetMe(response http.ResponseWriter, request *http.Request)
}
