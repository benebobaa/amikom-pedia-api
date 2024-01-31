package user_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/module/image/image_service"
	"amikom-pedia-api/module/user/user_service"
	"amikom-pedia-api/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type UserControllerImpl struct {
	UserService  user_service.UserService
	ImageService image_service.ImageService
}

func NewUserController(userService user_service.UserService, imageService image_service.ImageService) UserController {
	return &UserControllerImpl{UserService: userService, ImageService: imageService}
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
	// Check the Content-Type header to determine the request type
	contentType := request.Header.Get("Content-Type")

	authPayload := request.Header.Get(middleware.AuthorizationPayloadKey)
	userNId, _ := utils.FromStringToUsernameAndUUID(authPayload)
	userUUID := userNId.UserID

	if strings.Contains(contentType, "application/json") {
		// JSON request
		userUpdateRequest := user.UpdateRequestUser{}
		helper.ReadFromRequestBody(request, &userUpdateRequest)

		//authPayload := request.Header.Get(middleware.AuthorizationPayloadKey)
		//userNId, _ := utils.FromStringToUsernameAndUUID(authPayload)
		//userUUID := userNId.UserID

		userResponse := userController.UserService.Update(request.Context(), userUUID, userUpdateRequest)

		baseResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   userResponse,
		}

		helper.WriteToResponseBody(writer, baseResponse)
	} else if strings.Contains(contentType, "multipart/form-data") {
		// Form data request
		err := request.ParseMultipartForm(10 << 20) // 10 MB limit, adjust as needed
		if err != nil {
			http.Error(writer, "Unable to parse form data", http.StatusBadRequest)
			return
		}

		//Form-Data Request
		userUpdateRequest := user.UpdateRequestUser{
			Username: request.FormValue("username"),
			Name:     request.FormValue("name"),
			Bio:      request.FormValue("bio"),
			// Add other form fields as needed
		}
		_, imgHeaderhdr, err := request.FormFile("img_header")
		_, imgHeaderavtr, err := request.FormFile("img_avatar")
		if imgHeaderhdr != nil || imgHeaderavtr != nil {

			userController.ImageService.UploadToS3(request.Context(), userUUID, imgHeaderavtr, imgHeaderhdr)

		}

		//fmt.Println("imgHeaderhdr:", imgHeaderhdr)
		//
		//fmt.Println("imgHeaderavtr:", imgHeaderavtr)

		userResponse := userController.UserService.Update(request.Context(), userUUID, userUpdateRequest)

		baseResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   userResponse,
		}

		helper.WriteToResponseBody(writer, baseResponse)
	} else {
		// Unsupported content type
		http.Error(writer, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
	}
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

	authPayload := request.Header.Get(middleware.AuthorizationPayloadKey)

	userNId, _ := utils.FromStringToUsernameAndUUID(authPayload)

	userUUID := userNId.UserID

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
