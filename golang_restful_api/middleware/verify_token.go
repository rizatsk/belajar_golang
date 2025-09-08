package middleware

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/api"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func VerifyTokenMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		token := request.Header.Get("X-API-KEY")
		if token != "RAHASIATOKENRZ" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			apiResponse := api.ApiResponseError{
				Status:  "fail",
				Message: "Unathorized",
				Code:    "04010",
			}

			helper.WriteToResponseBody(writer, apiResponse)
		} else {
			next(writer, request, param)
		}
	}
}
