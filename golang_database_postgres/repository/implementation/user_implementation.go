package implementation

import (
	"context"
	"fmt"
	"golang_database_postgres/entity"
	"golang_database_postgres/repository/contract"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepositoryImplementation struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) contract.UserRepository {
	return &userRepositoryImplementation{DB: db}
}

func (repository *userRepositoryImplementation) Insert(context context.Context, user entity.User) (entity.User, error) {
	idV6, _ := uuid.NewV6()
	err := repository.DB.QueryRow(context, "INSERT INTO users (id, name, email, created_dt) VALUES($1, $2, $3, $4) RETURNING *",
		idV6.String(), user.Name, user.Email, time.Now().Format("2006-01-02 15:04:05")).
		Scan(&user.Id, &user.Email, &user.Name, &user.Created_dt)
	if err != nil {
		fmt.Println("QueryRow failed: \n", err)
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImplementation) FindAll(context context.Context) ([]entity.User, error) {
	rows, err_query := repository.DB.Query(
		context,
		"SELECT * FROM users ORDER BY created_dt DESC",
	)
	if err_query != nil {
		fmt.Println("Query failed: \n", err_query)
		return nil, err_query
	}
	defer rows.Close()

	// Mapping data users
	var allUser []entity.User
	for rows.Next() {
		var user_db entity.User
		err := rows.Scan(&user_db.Id, &user_db.Name, &user_db.Email, &user_db.Created_dt)
		if err != nil {
			panic(err)
		}

		allUser = append(allUser, user_db)
	}

	return allUser, err_query
}
