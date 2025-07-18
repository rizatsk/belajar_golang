package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Rizat")
	data.PushBack("Sakmir")
	data.PushBack("Software Engineer")

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
