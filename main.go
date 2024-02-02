package main

import (
	"fmt"

	"amikom-pedia-api/app"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	helper.PanicIfError(err)

	err = app.Serve(config)
	helper.PanicIfError(err)

	fmt.Println("Starting web server at port 3000")
}
