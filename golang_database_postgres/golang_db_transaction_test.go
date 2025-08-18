package golang_database_postgres

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDatabaseTransaction(test *testing.T) {
	// Connect to DB
	context := context.Background()
	pool_db := ConnectDatabasePool(context)
	defer pool_db.Close()

	// Create transaction
	db_trx, err_db_trx := pool_db.Begin(context)
	if err_db_trx != nil {
		fmt.Println("Failed create transaction:", err_db_trx)
		panic(err_db_trx)
	}
	// Rollback when have error
	defer db_trx.Rollback(context)

	user_rows := [][]interface{}{
		{GenerateUuidV6(), "Sangga2", "sangga2@gmail.com", time.Now().Format("2006-01-02 15:04:05")},
		{GenerateUuidV6(), "Agung2", "agung2@gmail.com", time.Now().Format("2006-01-02 15:04:05")},
		{GenerateUuidV6(), "Riko", "riko@gmail.com", time.Now().Format("2006-01-02 15:04:05")},
	}
	for _, row := range user_rows {
		_, err_exec := db_trx.Exec(context,
			"INSERT INTO users (id, name, email, created_dt) VALUES($1, $2, $3, $4)",
			row[0], row[1], row[2], row[3],
		)
		if err_exec != nil {
			fmt.Println("Exec failed: \n", err_exec)
		} else {
			fmt.Println("Success insert query exec")
		}
	}

	// commit transaction
	err_trx := db_trx.Commit(context)
	if err_trx != nil {
		fmt.Println("Failed commit transaction:", err_trx)
	}

	// Get all data users from db
	GetAllUserFromDB(context, pool_db)
}
