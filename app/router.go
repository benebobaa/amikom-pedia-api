package app

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/utils"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(config utils.Config) *httprouter.Router {
	controller := NewController(config)
	router := httprouter.New()
	midWare := middleware.NewMiddleware(router, controller.TokenMaker)

	// Exclude Auth Middleware
	router.POST("/api/v1/login", midWare.LoggingMiddleware(controller.LoginController.Login))
	router.POST("/api/v1/register", midWare.LoggingMiddleware(controller.RegisterController.Create))
	router.POST("/api/v1/users/forgot-password", midWare.LoggingMiddleware(controller.UserController.ForgotPassword))
	router.POST("/api/v1/otp/validate", midWare.LoggingMiddleware(controller.OTPController.Validation))
	// router.POST("/api/v1/otp/send", midWare.LoggingMiddleware(otpController.SendOtp))
	router.POST("/api/v1/otp/resend", midWare.LoggingMiddleware(controller.OTPController.ResendOtp))

	// Include Auth Middleware
	router.POST("/api/v1/users", controller.UserController.Create)
	router.GET("/api/v1/users", controller.UserController.FindAll)
	router.PUT("/api/v1/users/update", midWare.WrapperMiddleware(controller.UserController.Update))
	router.PUT("/api/v1/users/set-new-password", controller.UserController.SetNewPassword)
	router.PUT("/api/v1/users/change-password", midWare.WrapperMiddleware(controller.UserController.SetNewPassword))
	router.GET("/api/v1/users/:uuid", controller.UserController.FindByUUID)
	router.DELETE("/api/v1/users/:uuid", controller.UserController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
