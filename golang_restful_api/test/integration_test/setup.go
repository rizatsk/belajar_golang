package integrationtest

import (
	"context"
	"fmt"
	"golang_restful_api/api/v1/categories/controller"
	"golang_restful_api/app"
	"golang_restful_api/config"
	"golang_restful_api/exception"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"
	"golang_restful_api/repository"
	"golang_restful_api/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func setupTestDB() *pgxpool.Pool {
	// Load ENV
	err_load_env := godotenv.Load("../../.env.test")
	if err_load_env != nil {
		fmt.Println("Fail load file .env")
		panic(err_load_env)
	}

	context := context.Background()
	db := config.ConnectDatabasePool(context)

	return db
}

func setupTestRouter(db *pgxpool.Pool) http.Handler {
	validate := validator.New()

	// Dependency Injection
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(&app.Bootstrap{
		CategoryController: categoryController,
	})

	// Logging
	helper.LoggerInit()

	// Handle Error
	router.PanicHandler = exception.ErrorHandler
	router.NotFound = http.HandlerFunc(exception.NotFoundApiError)

	return router
}

func truncateCategory(context context.Context, db *pgxpool.Pool) {
	db.Query(context, "TRUNCATE TABLE categories;")
}

func createCategory(context context.Context, db *pgxpool.Pool, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(id, name) VALUES ($1, $2) RETURNING *"

	var respCategory domain.Category
	err := db.QueryRow(context, SQL,
		category.Id, category.Name).
		Scan(&respCategory.Id, &respCategory.Name)
	helper.PanicIfError(helper.PanicErrorParam{
		Err:     err,
		Ctx:     context,
		Message: "Error insert categories",
	})

	return category
}
