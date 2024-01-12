package main

import (
	"amikom-pedia-api/app"
	"amikom-pedia-api/helper"
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
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	config, err := utils.LoadConfig(".")
	helper.PanicIfError(err)

	gmailSender := mail.NewGmailSender(config.EmailName, config.EmailSender, config.EmailPassword)
	db := app.NewDB(config.DBDriver, config.DBSource)
	validate := validator.New()
	userRepository := user_repository.NewUserRepository()
	userService := user_service.NewUserService(userRepository, db, validate)
	userController := user_controller.NewUserController(userService)

	registerRepository := register_repository.NewRegisterRepository()
	otpRepository := otp_repository.NewOtpRepository()
	registerService := register_service.NewRegisterService(registerRepository, otpRepository, db, validate)
	registerController := register_controller.NewRegisterController(registerService)

	otpService := otp_service.NewOtpService(otpRepository, registerRepository, gmailSender, db, validate)
	otpController := otp_controller.NewOtpController(otpService)

	router := app.NewRouter(userController, registerController, otpController)

	server := http.Server{
		Addr:    config.ServerAddress,
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Starting web server at port 3000")
}
