package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/word.go"))
	fmt.Println(filepath.Base("hello/word.go"))
	fmt.Println(filepath.Ext("hello/word.go"))
	fmt.Println(filepath.Join("dodo", "hengki", "main.go"))
}
