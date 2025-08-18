package golang_database_postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_database_postgres/entity"
	"golang_database_postgres/repository/implementation"
	"testing"
)

func TestRepositoryInsert(test *testing.T) {
	context := context.Background()
	userRepository := implementation.NewUserRepository(ConnectDatabasePool(context))

	user := entity.User{
		Name:  "Rizat Test Repository",
		Email: "rizattestrepository@gmail.com",
	}
	result, err := userRepository.Insert(context, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestRepositoryFindAll(test *testing.T) {
	context := context.Background()
	userRepository := implementation.NewUserRepository(ConnectDatabasePool(context))

	result, err := userRepository.FindAll(context)
	if err != nil {
		panic(err)
	}

	// To Data JSON
	jsonSlice, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error marshalling slice:", err)
	}
	fmt.Println("Data all user:", string(jsonSlice))
}
