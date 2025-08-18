package golang_database_postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"golang_database_postgres/entity"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDatabase(context context.Context) (DB *pgx.Conn) {
	// Load ENV
	err_load_env := godotenv.Load()
	if err_load_env != nil {
		fmt.Println("Fail load file .env")
		panic(err_load_env)
	}

	connection_db, error_db := pgx.Connect(context, os.Getenv("DATABASE_URL"))
	if error_db != nil {
		fmt.Println("Connection database is fail : \n", error_db)
		panic(error_db)
	}

	fmt.Println("Success connection in DB")
	return connection_db
}

func TestGolangDatabasePostgres(test *testing.T) {
	// Connect to DB
	context := context.Background()
	db := ConnectDatabase(context)
	defer db.Close(context)

	// Query ke DB
	var user entity.User
	err := db.QueryRow(context, "SELECT id, name, email, created_dt FROM users WHERE email=$1", "rizatsakmir@gmail.com").Scan(&user.Id, &user.Email, &user.Name, &user.Created_dt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return
	}

	fmt.Printf("Data user: %+v\n", user)
}
