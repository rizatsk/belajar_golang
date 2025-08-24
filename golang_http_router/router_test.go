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

func TestRouter(test *testing.T) {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello http router")
	})

	request := httptest.NewRequest("GET", "http://localhost:6000/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(test, "Hello http router", string(bytes))
}

func TestRouterParam(test *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		data := "Product " + params.ByName("id")
		fmt.Fprint(writer, data)
	})

	request := httptest.NewRequest("GET", "http://localhost:6000/products/uuid-01", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(test, "Product uuid-01", string(bytes))
}

func TestRouterMultipleParam(test *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id/:name_user", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Println(params.ByName("id"))
		fmt.Println(params.ByName("name_user"))
		data := "Product " + params.ByName("id") + ": " + params.ByName("name_user")
		fmt.Fprint(writer, data)
	})

	request := httptest.NewRequest("GET", "http://localhost:6000/products/uuid-01/rizat", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(test, "Product uuid-01: rizat", string(bytes))
}

func TestRouterCatchAllParam(test *testing.T) {
	router := httprouter.New()

	router.GET("/images/*path", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Println(params.ByName("path"))
		data := "Is path image: " + params.ByName("path")
		fmt.Fprint(writer, data)
	})

	request := httptest.NewRequest("GET", "http://localhost:6000/images/user/123.webp", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(test, "Is path image: /user/123.webp", string(bytes))
}
