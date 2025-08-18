package golang_database_postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_database_postgres/entity"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllUserFromDB(context context.Context, pool_db *pgxpool.Pool) {
	// Get all users
	rows, err_query := pool_db.Query(
		context,
		"SELECT * FROM users ORDER BY created_dt DESC",
	)
	if err_query != nil {
		fmt.Println("Query failed: \n", err_query)
		return
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

	// To Data JSON
	jsonSlice, err := json.Marshal(allUser)
	if err != nil {
		fmt.Println("Error marshalling slice:", err)
		return
	}
	fmt.Println("Data all user:", string(jsonSlice))
}

func TestGolangQuerySql(test *testing.T) {
	// Connect to DB
	context := context.Background()
	pool_db := ConnectDatabasePool(context)
	defer pool_db.Close()

	// Query Insert
	var user entity.User
	idV6, _ := uuid.NewV6()
	err := pool_db.QueryRow(context, "INSERT INTO users (id, name, email, created_dt) VALUES($1, $2, $3, $4) RETURNING *",
		idV6.String(), "Jika", "jika@gmail.com", time.Now().Format("2006-01-02 15:04:05")).
		Scan(&user.Id, &user.Email, &user.Name, &user.Created_dt)
	if err != nil {
		fmt.Println("QueryRow failed: \n", err)
	} else {
		// Hasil user baru
		fmt.Println("Data new user :", user)
	}

	// Query Insert No Result
	id_user_exec, _ := uuid.NewV6()
	_, err_exec := pool_db.Exec(context,
		"INSERT INTO users (id, name, email, created_dt) VALUES($1, $2, $3, $4)",
		id_user_exec.String(), "Joko", "joko@gmail.com", time.Now().Format("2006-01-02 15:04:05"))
	if err_exec != nil {
		fmt.Println("Exec failed: \n", err_exec)
	} else {
		fmt.Println("Success insert query exec")
	}

	GetAllUserFromDB(context, pool_db)
}
