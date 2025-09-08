package exception

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/api"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, err) {
		return
	}

	if validationError(writer, err) {
		return
	}

	internalServerError(writer, err)
}

func NotFoundApiError(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)

	apiResponse := api.ApiResponseError{
		Status:  "fail",
		Message: "Route is not found",
		Code:    "04044",
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func validationError(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponse := api.ApiResponseError{
			Status:  "fail",
			Message: "Bad request",
			Code:    "04000",
			Error:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponse := api.ApiResponseError{
			Status:  "fail",
			Message: "Data is not found",
			Code:    "04040",
			Error:   exception.Error,
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	apiResponse := api.ApiResponseError{
		Status:  "fail",
		Message: "Internal server error",
		Code:    "05000",
		Error:   err,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
