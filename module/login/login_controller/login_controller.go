package login_controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type LoginController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
