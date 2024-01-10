package user_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/module/user/user_service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	UserService user_service.UserService
}

func NewUserController(userService user_service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (userController *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := user.CreateRequestUser{}

	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := userController.UserService.Create(request.Context(), userCreateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
