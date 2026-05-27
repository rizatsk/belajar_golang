package helper

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load ENV
	err_load_env := godotenv.Load()
	if err_load_env != nil {
		fmt.Println("Fail load file .env")
		panic(err_load_env)
	}
}
