package otp_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/otp"
	"amikom-pedia-api/module/otp/otp_service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OtpControllerImpl struct {
	OtpService otp_service.OtpService
}

func NewOtpController(otpService otp_service.OtpService) OtpController {
	return &OtpControllerImpl{OtpService: otpService}
}

func (otpController *OtpControllerImpl) Validation(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	otpValidateRequest := otp.OtpValidateRequest{}
	helper.ReadFromRequestBody(request, &otpValidateRequest)

	otpController.OtpService.Validation(request.Context(), otpValidateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
