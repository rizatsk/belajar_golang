package controller

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/api"
	"golang_restful_api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := api.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	apiResponse := api.ApiResponse{
		Status:  "success",
		Message: "Success create category",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := api.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryResponse := controller.CategoryService.Upate(request.Context(), categoryUpdateRequest)
	apiResponse := api.ApiResponse{
		Status:  "success",
		Message: "Success update category",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	category_id := params.ByName("categoryId")

	controller.CategoryService.Delete(request.Context(), category_id)
	apiResponse := api.ApiResponse{
		Status:  "success",
		Message: "Success update category",
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	category_id := params.ByName("categoryId")

	categoryResponse := controller.CategoryService.FindById(request.Context(), category_id)
	apiResponse := api.ApiResponse{
		Status:  "success",
		Message: "Success get category",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(request.Context())

	apiResponse := api.ApiResponse{
		Status:  "success",
		Message: "Success get all category",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
