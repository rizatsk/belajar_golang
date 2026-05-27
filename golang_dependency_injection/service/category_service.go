package service

import (
	"context"
	"golang_restful_api/exception"
	"golang_restful_api/helper"
	"golang_restful_api/model/api"
	"golang_restful_api/model/domain"
	"golang_restful_api/repository"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryService interface {
	Create(ctx context.Context, request api.CategoryCreateRequest) api.CategoryResponse
	Upate(ctx context.Context, request api.CategoryUpdateRequest) api.CategoryResponse
	Delete(ctx context.Context, category_id string)
	FindById(ctx context.Context, category_id string) api.CategoryResponse
	FindAll(ctx context.Context) []api.CategoryResponse
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *pgxpool.Pool
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *pgxpool.Pool, Validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           Validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request api.CategoryCreateRequest) api.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(helper.PanicErrorParam{
		Err:     err,
		Ctx:     ctx,
		Message: "Error validation when create category",
	})

	category := domain.Category{
		Id:   helper.GenerateUuidV6(),
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, service.DB, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Upate(ctx context.Context, request api.CategoryUpdateRequest) api.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(helper.PanicErrorParam{
		Err:     err,
		Ctx:     ctx,
		Message: "Error validation when update category",
	})

	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category = service.CategoryRepository.Update(ctx, service.DB, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, category_id string) {
	rowAffected := service.CategoryRepository.Delete(ctx, service.DB, category_id)

	if rowAffected < 1 {
		panic(exception.NewNotFoundError("Data is not found", ctx))
	}
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, category_id string) api.CategoryResponse {
	category, err := service.CategoryRepository.FindById(ctx, service.DB, category_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error(), ctx))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []api.CategoryResponse {
	category := service.CategoryRepository.FindAll(ctx, service.DB)

	return helper.ToCategoryResponses(category)
}
