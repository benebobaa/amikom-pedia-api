package user_controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUUID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SetNewPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
