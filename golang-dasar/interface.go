package main

import "fmt"

type HashName interface {
	GetName() string
}

type Person struct {
	Name string
}

func sayHello(value HashName) {
	fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{Name: "Bobby"}
	sayHello(person)
}
