package main

import (
	"amikom-pedia-api/app"
	"amikom-pedia-api/controller"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/repository"
	"amikom-pedia-api/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewUserRepository()
	categoryService := service.NewUserService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Starting web server at port 3000")
}
