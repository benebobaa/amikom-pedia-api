package user_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/module/user/user_service"
	"amikom-pedia-api/utils"
	"fmt"
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

func (userController *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := user.UpdateRequestUser{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	authPayload := request.Header.Get(middleware.AuthorizationPayloadKey)

	userNId, _ := utils.FromStringToUsernameAndUUID(authPayload)

	fmt.Println("userNId:", userNId)
	fmt.Println("USER UUID:", userNId.UserID)
	userUUID := userNId.UserID

	userResponse := userController.UserService.Update(request.Context(), userUUID, userUpdateRequest)

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

func (userController *UserControllerImpl) ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userForgotPassword := user.ForgotPasswordRequest{}

	helper.ReadFromRequestBody(request, &userForgotPassword)

	result := userController.UserService.ForgotPassword(request.Context(), userForgotPassword.Email)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController *UserControllerImpl) SetNewPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userSetNewPassword := user.SetNewPasswordRequest{}

	helper.ReadFromRequestBody(request, &userSetNewPassword)

	userController.UserService.SetNewPassword(request.Context(), userSetNewPassword)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController *UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	newPasswordRequest := user.UpdatePasswordRequest{}
	helper.ReadFromRequestBody(request, &newPasswordRequest)

	userUUID := newPasswordRequest.UUID

	err := userController.UserService.UpdatePassword(request.Context(), userUUID, newPasswordRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
