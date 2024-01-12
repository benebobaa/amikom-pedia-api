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

func (userController *UserControllerImpl) FindByUUID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	uuid := params.ByName("uuid")

	userResponse := userController.UserService.FindByUUID(request.Context(), uuid)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	usersResponse := userController.UserService.FindAll(request.Context())

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   usersResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	uuid := params.ByName("uuid")

	userController.UserService.Delete(request.Context(), uuid)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}