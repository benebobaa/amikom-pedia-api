package app

import (
	"net/http"

	"amikom-pedia-api/utils"
)

func Serve(config utils.Config) error {
	router := NewRouter(config)
	server := http.Server{
		Addr:    config.ServerAddress,
		Handler: router,
	}

	return server.ListenAndServe()
}
