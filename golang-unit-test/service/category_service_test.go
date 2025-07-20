package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(test *testing.T) {
	test.Run("Category no have data", func(test *testing.T) {
		// Arrange
		categoryRepository.Mock.On("FindById", "1").Return(nil)

		// Act
		result, err := categoryService.Get("1")

		// Assert
		assert.Nil(test, result, "Result must be nil")
		assert.NotNil(test, err, "Error must be not nil")
	})
	test.Run("Category have data", func(test *testing.T) {
		// Arrange
		category := entity.Category{
			Id:   "2",
			Name: "Laptop",
		}
		categoryRepository.Mock.On("FindById", "2").Return(category)

		// Act
		result, err := categoryService.Get("2")

		// Assert
		assert.Nil(test, err, "Error must be nil")
		assert.Equal(test, category, *result, "Result is not equal category")
	})
}
