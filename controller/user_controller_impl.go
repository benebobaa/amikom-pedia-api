package controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewCategoryController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (userControllerImpl *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := user.CreateRequestUser{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	categoryResponse := userControllerImpl.UserService.Create(request.Context(), userCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (userControllerImpl *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequestPassword := user.CreateUpdatePassword{}
	helper.ReadFromRequestBody(request, &userRequestPassword)

	userId := params.ByName("userId")
	//id, err := strconv.Atoi(categoryId)
	//helper.PanicIfError(err)

	userRequestPassword.Id = userId

	categoryResponse := userControllerImpl.UserService.Update(request.Context(), userRequestPassword)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (userControllerImpl *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	//id, err := strconv.Atoi(categoryId)
	//helper.PanicIfError(err)

	userControllerImpl.UserService.Delete(request.Context(), userId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (userControllerImpl *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	//id, err := strconv.Atoi(categoryId)
	//helper.PanicIfError(err)

	categoryResponse := userControllerImpl.UserService.FindById(request.Context(), userId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (userControllerImpl *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := userControllerImpl.UserService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
