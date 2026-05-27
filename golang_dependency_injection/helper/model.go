package helper

import (
	"golang_restful_api/model/api"
	"golang_restful_api/model/domain"
)

func ToCategoryResponse(category domain.Category) api.CategoryResponse {
	return api.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []api.CategoryResponse {
	var categoryResponses []api.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
