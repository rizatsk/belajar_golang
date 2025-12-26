package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"
	"golang_restful_api/model/logger"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategorySuccess(t *testing.T) {
	context := context.Background()
	db := setupTestDB()

	truncateCategory(context, db)
	router := setupTestRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIATOKENRZ")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	helper.LoggerDebug(logger.LoggerDebug{
		Message: "Result data create category success",
		Data:    responseBody,
	})

	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "Success create category", responseBody["message"])
	assert.NotEmpty(t, responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFail401(t *testing.T) {
	context := context.Background()
	db := setupTestDB()

	truncateCategory(context, db)
	router := setupTestRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	helper.LoggerDebug(logger.LoggerDebug{
		Message: "Result data create category fail 401",
		Data:    responseBody,
	})

	assert.Equal(t, "fail", responseBody["status"])
	assert.Equal(t, "04010", responseBody["code"])
	assert.Equal(t, "Unathorized", responseBody["message"])
	assert.Nil(t, responseBody["error"])
}

func TestPatchCategorySuccess(t *testing.T) {
	context := context.Background()
	db := setupTestDB()

	truncateCategory(context, db)

	category := domain.Category{
		Id:   helper.GenerateUuidV6(),
		Name: "Mobil",
	}
	createCategory(context, db, category)
	router := setupTestRouter(db)

	requestBody := strings.NewReader(fmt.Sprintf(`{"id": "%s", "name": "Motor"}`, category.Id))
	request := httptest.NewRequest(http.MethodPatch, "http://localhost:3000/api/v1/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIATOKENRZ")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	helper.LoggerDebug(logger.LoggerDebug{
		Message: "Result data patch category success",
		Data:    responseBody,
	})

	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "Success update category", responseBody["message"])
	assert.Equal(t, category.Id, responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, "Motor", responseBody["data"].(map[string]interface{})["name"])
}
