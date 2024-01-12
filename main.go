package main

import (
	"amikom-pedia-api/app"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/module/login/login_controller"
	"amikom-pedia-api/module/login/login_repository"
	"amikom-pedia-api/module/login/login_service"
	"amikom-pedia-api/module/register/register_controller"
	"amikom-pedia-api/module/register/register_repository"
	"amikom-pedia-api/module/register/register_service"
	"amikom-pedia-api/module/user/user_controller"
	"amikom-pedia-api/module/user/user_repository"
	"amikom-pedia-api/module/user/user_service"
	"amikom-pedia-api/utils"
	"amikom-pedia-api/utils/token"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	config, err := utils.LoadConfig(".")
	helper.PanicIfError(err)

	//jwt init
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	helper.PanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	validate := validator.New()
	userRepository := user_repository.NewUserRepository()
	userService := user_service.NewUserService(userRepository, db, validate)
	userController := user_controller.NewUserController(userService)

	registerRepository := register_repository.NewRegisterRepository()
	registerService := register_service.NewRegisterService(registerRepository, db, validate)
	registerController := register_controller.NewRegisterController(registerService)

	//LOGIN INTERFACE
	loginRepository := login_repository.NewLoginRepository()
	loginService := login_service.NewLoginService(tokenMaker, loginRepository, db, validate)
	loginController := login_controller.NewLoginController(loginService)

	router := app.NewRouter(userController, registerController, loginController)

	server := http.Server{
		Addr:    config.ServerAddress,
		Handler: middleware.NewAuthMiddleware(router, tokenMaker),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Starting web server at port 3000")
}
