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

type LogMidlleware struct {
	http.Handler
}

func (midleware *LogMidlleware) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive request")
	// Akan diteruskan ke handler nya
	midleware.Handler.ServeHTTP(writter, request)
}

func TestMidleware(test *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello http router")
	})

	middleware := LogMidlleware{router}

	request := httptest.NewRequest("GET", "http://localhost:6000/", nil)
	recoder := httptest.NewRecorder()

	middleware.ServeHTTP(recoder, request)
	response := recoder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(test, "Hello http router", string(bytes))
}
