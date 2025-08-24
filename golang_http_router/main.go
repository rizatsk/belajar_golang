package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed public
var public embed.FS

func main() {
	router := httprouter.New()
	directory, _ := fs.Sub(public, "public")
	router.ServeFiles("/public/*filepath", http.FS(directory))
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello http router")
	})

	addr := "localhost:6000"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	fmt.Println("Serve running", addr)
	server.ListenAndServe()
}
