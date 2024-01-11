package register_controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RegisterController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
