package register_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/register"
	"amikom-pedia-api/module/register/register_service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RegisterControllerImpl struct {
	RegisterService register_service.RegisterService
}

func NewRegisterController(registerService register_service.RegisterService) RegisterController {
	return &RegisterControllerImpl{RegisterService: registerService}
}

func (registerController *RegisterControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerCreateRequest := register.RegisterRequest{}

	helper.ReadFromRequestBody(request, &registerCreateRequest)

	registerResponse := registerController.RegisterService.Create(request.Context(), registerCreateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   registerResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
