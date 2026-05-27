package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabasePool(context context.Context) (DB *pgxpool.Pool) {
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
