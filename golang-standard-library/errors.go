package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("Is validation error")
	notFoundError   = errors.New("Is not found error")
)

func cekName(name string) error {
	if name == "" {
		return validationError
	} else if name != "Agus" {
		return notFoundError
	}

	return nil
}

func main() {
	errorCekName := cekName("Rizat")
	if errorCekName != nil {
		if errors.Is(errorCekName, validationError) {
			fmt.Println("Validation error")
		} else if errors.Is(errorCekName, notFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unkown error")
		}
	}
}
