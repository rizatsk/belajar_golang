package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		log.Printf("[%s] %s", request.Method, request.URL.Path)
		next(writer, request, param)
	}
}
