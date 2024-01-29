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

	otpResponse := otpController.OtpService.Validation(request.Context(), otpValidateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   otpResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (otpController *OtpControllerImpl) SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createRequestOtp := otp.SendOtpRequest{}
	helper.ReadFromRequestBody(request, &createRequestOtp)

	err := otpController.OtpService.SendOtp(request.Context(), createRequestOtp)
	helper.PanicIfError(err)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
