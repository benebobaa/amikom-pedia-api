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

	/// Exclude Auth Middleware
	//AUTH
	router.POST("/api/v1/login", midWare.LoggingMiddleware(controller.LoginController.Login))
	router.POST("/api/v1/register", midWare.LoggingMiddleware(controller.RegisterController.Create))

	//OTP
	router.POST("/api/v1/users/forgot-password", midWare.LoggingMiddleware(controller.UserController.ForgotPassword))
	router.POST("/api/v1/otp/validate", midWare.LoggingMiddleware(controller.OTPController.Validation))
	// router.POST("/api/v1/otp/send", midWare.LoggingMiddleware(otpController.SendOtp))
	router.POST("/api/v1/otp/resend", midWare.LoggingMiddleware(controller.OTPController.ResendOtp))

	/// Include Auth Middleware
	// USERS
	router.POST("/api/v1/users", controller.UserController.Create)
	router.GET("/api/v1/users", controller.UserController.FindAll)
	router.PUT("/api/v1/users/update", midWare.WrapperMiddleware(controller.UserController.Update))
	router.PUT("/api/v1/users/set-new-password", controller.UserController.SetNewPassword)
	router.PUT("/api/v1/users/change-password", midWare.WrapperMiddleware(controller.UserController.UpdatePassword))
	router.GET("/api/v1/users/profile", midWare.WrapperMiddleware(controller.UserController.FindByUUID))
	router.DELETE("/api/v1/users/:uuid", controller.UserController.Delete)

	//POSTS
	router.POST("/api/v1/post", midWare.WrapperMiddleware(controller.PostController.Create))
	router.PUT("/api/v1/post/:id", midWare.WrapperMiddleware(controller.PostController.Update))
	router.GET("/api/v1/post", midWare.WrapperMiddleware(controller.PostController.FindAll))
	router.GET("/api/v1/post/:id", midWare.WrapperMiddleware(controller.PostController.FindById))
	router.DELETE("/api/v1/post/:id", midWare.WrapperMiddleware(controller.PostController.Delete))

	router.PanicHandler = exception.ErrorHandler

	return router
}
