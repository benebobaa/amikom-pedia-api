package login_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/login"
	"amikom-pedia-api/module/login/login_service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type LoginControllerImpl struct {
	LoginService login_service.LoginService
}

func NewLoginController(loginService login_service.LoginService) LoginController {
	return &LoginControllerImpl{LoginService: loginService}
}

func (loginController *LoginControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginCreateRequest := login.LoginRequest{}

	helper.ReadFromRequestBody(request, &loginCreateRequest)

	registerResponse := loginController.LoginService.Login(request.Context(), loginCreateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   registerResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
