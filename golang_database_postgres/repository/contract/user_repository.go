package contract

import (
	"context"
	"golang_database_postgres/entity"
)

type UserRepository interface {
	Insert(context context.Context, user entity.User) (entity.User, error)
	FindAll(context context.Context) ([]entity.User, error)
}
