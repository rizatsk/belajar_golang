package golang_database_postgres

import (
	"context"
	"fmt"
	"golang_database_postgres/entity"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDatabasePool(context context.Context) (DB *pgxpool.Pool) {
	// Load ENV
	err_load_env := godotenv.Load()
	if err_load_env != nil {
		fmt.Println("Fail load file .env")
		panic(err_load_env)
	}

	// Open connection
	connection_db, error_db := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if error_db != nil {
		fmt.Println("Connection database is fail : \n", error_db)
		panic(error_db)
	}

	// Contoh atur max connection
	connection_db.MaxConns = 50
	connection_db.MinConns = 10
	connection_db.MaxConnLifetime = time.Hour
	connection_db.MaxConnIdleTime = 30 * time.Minute
	fmt.Println("Success connection in DB")

	// Buat pool
	pool_db, err_pool_db := pgxpool.NewWithConfig(context, connection_db)
	if err_pool_db != nil {
		fmt.Println("Fail create pool : \n", err_pool_db)
		panic(err_pool_db)
	}

	return pool_db
}

func TestGolangDatabasePostgresPool(test *testing.T) {
	// Load ENV
	err_load_env := godotenv.Load()
	if err_load_env != nil {
		fmt.Println("Fail load file .env")
		panic(err_load_env)
	}

	context := context.Background()
	// Connect to DB
	pool_db := ConnectDatabasePool(context)
	defer pool_db.Close()

	// Query ke DB
	var user entity.User
	err := pool_db.QueryRow(context, "SELECT id, name, email, created_dt FROM users WHERE email=$1", "rizatsakmir@gmail.com").Scan(&user.Id, &user.Email, &user.Name, &user.Created_dt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return
	}

	fmt.Printf("Data user: %+v\n", user)
}
