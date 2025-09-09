package main

import (
	"fmt"
	"golang_restful_api/app"
	"golang_restful_api/config"
	"golang_restful_api/exception"
	"golang_restful_api/helper"
	"net/http"
)

func main() {
	bootsrap := app.InitApp()
	router := app.NewRouter(bootsrap)

	// Logging
	config.LoggerInit()

	// Handle Error
	router.PanicHandler = exception.ErrorHandler
	router.NotFound = http.HandlerFunc(exception.NotFoundApiError)

	addr := "localhost:6000"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	fmt.Println("Serve running", addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
