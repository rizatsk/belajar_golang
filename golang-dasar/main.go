package main

import (
	"fmt"
	"golang-dasar/utils"
)

func sayHello(name string, job string) {
	fmt.Println("Running function say hello")
	fmt.Println("Hello ", name)
	fmt.Println("Your jobs is ", job)
}

type User struct {
	ID   int
	Name string
}

func sayHelloObject(user User) {
	fmt.Println("Running function say hello object")
	fmt.Println("User id", user.ID)
	fmt.Println("Your name", user.Name)
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("This is number int: ", 1)
	fmt.Println("This is number float: ", 0.38)

	fmt.Println("This is true: ", true)
	fmt.Println("This is false: ", false)

	var name string

	name = "Hello world"
	fmt.Println(name)

	name = "Hello rizat"
	fmt.Println(name)

	var name_user = "Rizat Sakmir"
	fmt.Println(name_user)

	job_user := "Software Engineer"
	fmt.Println(job_user)

	var (
		firstName = "Rizat"
		lastName  = "Sakmir"
	)
	fmt.Println("Fullname : ", firstName, lastName)

	const user_job = "Programmer"
	fmt.Println(user_job)

	const convertString = "Rizat Sakmir"
	var s = string(convertString[6])
	fmt.Println(s)

	type NoKTP string

	var ktpRz NoKTP = "0001111"
	var ktpExample NoKTP = NoKTP(ktpRz)
	fmt.Println(ktpExample)

	var slice = []int{123, 124, 125}
	fmt.Println(slice)

	for counter := 1; counter <= 10; counter++ {
		if counter == 6 {
			break
		}

		if counter == 2 {
			continue
		}

		fmt.Println("Perulangan ke ", counter)
	}

	names := []string{"Bambang", "Joko", "Rizat"}
	for index, name := range names {
		fmt.Println("This is index: ", index)
		fmt.Println("This is name: ", name)
	}

	learnerSlice()
	sayHello("Rizat", "Software Engineer")
	sayHelloObject(User{ID: 1, Name: "Rizat"})

	utils.TestAccess()
}
