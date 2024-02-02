package app

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/module/login/login_controller"
	"amikom-pedia-api/module/otp/otp_controller"
	"amikom-pedia-api/module/post/post_controller"
	"amikom-pedia-api/module/register/register_controller"
	"amikom-pedia-api/module/user/user_controller"
	"amikom-pedia-api/utils/token"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(tokenMaker token.Maker, userController user_controller.UserController, registerController register_controller.RegisterController, otpController otp_controller.OtpController, loginController login_controller.LoginController, postController post_controller.PostController) *httprouter.Router {
	router := httprouter.New()
	midWare := middleware.NewMiddleware(router, tokenMaker)

	//Exclude Auth Middleware
	router.POST("/api/v1/login", midWare.LoggingMiddleware(loginController.Login))
	router.POST("/api/v1/register", midWare.LoggingMiddleware(registerController.Create))
	router.POST("/api/v1/users/forgot-password", midWare.LoggingMiddleware(userController.ForgotPassword))
	router.POST("/api/v1/otp/validate", midWare.LoggingMiddleware(otpController.Validation))
	//router.POST("/api/v1/otp/send", midWare.LoggingMiddleware(otpController.SendOtp))
	router.POST("/api/v1/otp/resend", midWare.LoggingMiddleware(otpController.ResendOtp))

	//Include Auth Middleware
	router.POST("/api/v1/users", userController.Create)
	router.GET("/api/v1/users", userController.FindAll)
	router.PUT("/api/v1/users/update", midWare.WrapperMiddleware(userController.Update))
	router.PUT("/api/v1/users/set-new-password", userController.SetNewPassword)
	router.PUT("/api/v1/users/change-password", midWare.WrapperMiddleware(userController.UpdatePassword))
	router.GET("/api/v1/users/:uuid", userController.FindByUUID)
	router.DELETE("/api/v1/users/:uuid", userController.Delete)
	router.POST("/api/v1/post", midWare.WrapperMiddleware(postController.Create))
	router.PUT("/api/v1/post/:id", midWare.WrapperMiddleware(postController.Update))
	router.GET("/api/v1/post", postController.FindAll)
	router.GET("/api/v1/post/:id", postController.FindById)
	router.DELETE("/api/v1/post/:id", postController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
