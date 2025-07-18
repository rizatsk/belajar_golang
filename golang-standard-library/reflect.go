package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name, Email, Phone string `required:"true"`
}

func isValid(value any) (result bool) {
	result = true
	dataType := reflect.TypeOf(value)
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	isValidUser := isValid(User{
		Name:  "Dodo",
		Email: "dodo@gmail.com",
		Phone: "081231412314",
	})
	fmt.Println("Validasi user :", isValidUser)
}
