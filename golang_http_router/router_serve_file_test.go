package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed public
var resources embed.FS

func TestServeFile(test *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "public")
	router.ServeFiles("/public/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:6000/public/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(test, "hello world", string(body))
}
