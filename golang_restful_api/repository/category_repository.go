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
	helper.PanicIfError(err)

	return respCategory
}

func (repository *CategoryRepositoryImpl) Update(context context.Context, pool_db *pgxpool.Pool, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = $2 WHERE id = $1"

	_, err := pool_db.Exec(context, SQL, category.Id, category.Name)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(context context.Context, pool_db *pgxpool.Pool, category_id string) int64 {
	SQL := "DELETE FROM categories WHERE id = $1"

	tag, err := pool_db.Exec(context, SQL, category_id)
	helper.PanicIfError(err)

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

		helper.PanicIfError(err)
	}

	return respCategory, nil
}

func (repository *CategoryRepositoryImpl) FindAll(context context.Context, pool_db *pgxpool.Pool) []domain.Category {
	SQL := "SELECT * FROM categories"

	rows, err := pool_db.Query(context, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var allCategories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		allCategories = append(allCategories, category)
	}

	return allCategories
}
