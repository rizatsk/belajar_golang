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

func TestRouterPanicHandler(test *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprint(writer, "Panic ", i)
	}
	router.GET("/panic-handler", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Internal server error")
	})

	request := httptest.NewRequest("GET", "http://localhost:6000/panic-handler", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(test, "Panic Internal server error", string(body))
}
