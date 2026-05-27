package repository

import (
	"context"
	"errors"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository interface {
	Save(ctx context.Context, pool_db *pgxpool.Pool, category domain.Category) domain.Category
	Update(ctx context.Context, pool_db *pgxpool.Pool, cateory domain.Category) domain.Category
	Delete(ctx context.Context, pool_db *pgxpool.Pool, category_id string) int64
	FindById(ctx context.Context, pool_db *pgxpool.Pool, category_id string) (domain.Category, error)
	FindAll(ctx context.Context, pool_db *pgxpool.Pool) []domain.Category
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(context context.Context, pool_db *pgxpool.Pool, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(id, name) VALUES ($1, $2) RETURNING *"

	var respCategory domain.Category
	err := pool_db.QueryRow(context, SQL,
		category.Id, category.Name).
		Scan(&respCategory.Id, &respCategory.Name)
	helper.PanicIfError(helper.PanicErrorParam{
		Err:     err,
		Ctx:     context,
		Message: "Error insert categories",
	})

	return respCategory
}

func (repository *CategoryRepositoryImpl) Update(context context.Context, pool_db *pgxpool.Pool, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = $2 WHERE id = $1"

	_, err := pool_db.Exec(context, SQL, category.Id, category.Name)
	helper.PanicIfError(helper.PanicErrorParam{
		Err:     err,
		Ctx:     context,
		Message: "Error update categories",
	})

	return category
}

func (repository *CategoryRepositoryImpl) Delete(context context.Context, pool_db *pgxpool.Pool, category_id string) int64 {
	SQL := "DELETE FROM categories WHERE id = $1"

	tag, err := pool_db.Exec(context, SQL, category_id)
	helper.PanicIfError(helper.PanicErrorParam{
		Err:     err,
		Ctx:     context,
		Message: "Error delete categories",
	})

	return tag.RowsAffected()
}

func (repository *CategoryRepositoryImpl) FindById(context context.Context, pool_db *pgxpool.Pool, category_id string) (domain.Category, error) {
	SQL := "SELECT * FROM categories WHERE id = $1 LIMIT 1"

	var respCategory domain.Category
	err := pool_db.QueryRow(context, SQL, category_id).
		Scan(&respCategory.Id, &respCategory.Name)
	if err == pgx.ErrNoRows {
		return respCategory, errors.New("category is not found")
	} else {
		helper.PanicIfError(helper.PanicErrorParam{
			Err:     err,
			Ctx:     context,
			Message: "Category is not found where find by id",
		})
	}

	return respCategory, nil
}

func (repository *CategoryRepositoryImpl) FindAll(context context.Context, pool_db *pgxpool.Pool) []domain.Category {
	SQL := "SELECT * FROM categories"

	rows, err := pool_db.Query(context, SQL)
	helper.PanicIfError(helper.PanicErrorParam{
		Err:     err,
		Ctx:     context,
		Message: "Error find all categories",
	})

	defer rows.Close()

	var allCategories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(helper.PanicErrorParam{
			Err:     err,
			Ctx:     context,
			Message: "Error when map data find all categories",
		})

		allCategories = append(allCategories, category)
	}

	return allCategories
}
