package main

import (
	"amikom-pedia-api/app"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/module/login/login_controller"
	"amikom-pedia-api/module/login/login_repository"
	"amikom-pedia-api/module/login/login_service"
	"amikom-pedia-api/module/otp/otp_controller"
	"amikom-pedia-api/module/otp/otp_repository"
	"amikom-pedia-api/module/otp/otp_service"
	"amikom-pedia-api/module/register/register_controller"
	"amikom-pedia-api/module/register/register_repository"
	"amikom-pedia-api/module/register/register_service"
	"amikom-pedia-api/module/user/user_controller"
	"amikom-pedia-api/module/user/user_repository"
	"amikom-pedia-api/module/user/user_service"
	"amikom-pedia-api/utils"
	"amikom-pedia-api/utils/mail"
	"amikom-pedia-api/utils/token"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	config, err := utils.LoadConfig(".")
	helper.PanicIfError(err)
	//jwt init
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	helper.PanicIfError(err)

	gmailSender := mail.NewGmailSender(config.EmailName, config.EmailSender, config.EmailPassword)
	db := app.NewDB(config.DBDriver, config.DBSource)
	validate := utils.CustomValidator()

	//REPOSITORY
	userRepository := user_repository.NewUserRepository()
	registerRepository := register_repository.NewRegisterRepository()
	otpRepository := otp_repository.NewOtpRepository()
	loginRepository := login_repository.NewLoginRepository()

	//SERVICE
	userService := user_service.NewUserService(userRepository, otpRepository, db, validate)
	registerService := register_service.NewRegisterService(registerRepository, otpRepository, db, validate)
	loginService := login_service.NewLoginService(tokenMaker, loginRepository, db, validate)
	//otpService := otp_service.NewOtpService(otpRepository, registerRepository, gmailSender, db, validate)
	otpService := otp_service.NewOtpService(otpRepository, registerRepository, userRepository, gmailSender, db, validate)

	//CONTROLLER
	userController := user_controller.NewUserController(userService)
	registerController := register_controller.NewRegisterController(registerService)
	loginController := login_controller.NewLoginController(loginService)
	otpController := otp_controller.NewOtpController(otpService)

	router := app.NewRouter(userController, registerController, otpController, loginController)

	server := http.Server{
		Addr:    config.ServerAddress,
		Handler: middleware.NewAuthMiddleware(router, tokenMaker),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Starting web server at port 3000")
}
