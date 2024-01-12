package otp_controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OtpController interface {
	Validation(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
