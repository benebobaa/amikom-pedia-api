package app

import (
	"amikom-pedia-api/controller"
	"amikom-pedia-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", categoryController.FindAll)
	router.GET("/api/users/:userId", categoryController.FindById)
	router.POST("/api/users", categoryController.Create)
	router.PUT("/api/users/:userId", categoryController.Update)
	router.DELETE("/api/users/:userId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
