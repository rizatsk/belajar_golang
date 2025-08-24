package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterMethodNotAllowedHandler(test *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Method is not allowed")
	})

	router.GET("/method-not-allowed", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hai i'm Iron Man")
	})

	request := httptest.NewRequest("POST", "http://localhost:6000/method-not-allowed", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(test, "Method is not allowed", string(body))
}
