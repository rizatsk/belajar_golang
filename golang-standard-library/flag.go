package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "Database username")
	password := flag.String("password", "root", "Database password")

	flag.Parse()

	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
}
