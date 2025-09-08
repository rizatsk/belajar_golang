package app

import (
	"context"
	"fmt"
	"golang_restful_api/api/v1/categories/controller"
	"golang_restful_api/config"
	"golang_restful_api/repository"
	"golang_restful_api/service"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

type Bootstrap struct {
	CategoryController controller.CategoryController
}

func InitApp() *Bootstrap {
	err_load_env := godotenv.Load()
	if err_load_env != nil {
		fmt.Println("Fail load file .env")
		panic(err_load_env)
	}

	context := context.Background()
	db := config.ConnectDatabasePool(context)
	validate := validator.New()

	// Dependency Injection
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	return &Bootstrap{
		CategoryController: categoryController,
	}
}
