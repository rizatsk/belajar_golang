package golang_database_postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

func GenerateUuidV6() string {
	uid, _ := uuid.NewV6()
	return uid.String()
}

func TestCopyFrom(test *testing.T) {
	// Connect to DB
	context := context.Background()
	pool_db := ConnectDatabasePool(context)
	defer pool_db.Close()

	user_rows := [][]interface{}{
		{GenerateUuidV6(), "Sangga", "sangga@gmail.com", time.Now().Format("2006-01-02 15:04:05")},
		{GenerateUuidV6(), "Agung", "agung@gmail.com", time.Now().Format("2006-01-02 15:04:05")},
		{GenerateUuidV6(), "Riko", "riko@gmail.com", time.Now().Format("2006-01-02 15:04:05")},
	}

	// Copy from for insert many data
	_, err := pool_db.CopyFrom(
		context,
		pgx.Identifier{"users"},
		[]string{"id", "name", "email", "created_dt"},
		pgx.CopyFromRows(user_rows),
	)
	if err != nil {
		fmt.Println("Error CopyFrom: ", err)
	}

	// Get all data users from db
	GetAllUserFromDB(context, pool_db)
}
