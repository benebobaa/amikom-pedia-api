package app

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/module/login/login_controller"
	"amikom-pedia-api/module/register/register_controller"
	"amikom-pedia-api/module/user/user_controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController user_controller.UserController, registerController register_controller.RegisterController, loginController login_controller.LoginController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/users", userController.Create)
	router.GET("/api/v1/users/:uuid", userController.FindByUUID)
	router.GET("/api/v1/users", userController.FindAll)
	router.DELETE("/api/v1/users/:uuid", userController.Delete)
	router.POST("/api/v1/register", registerController.Create)

	//LOGIN
	router.POST("/api/v1/login", loginController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
