package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666) // 0666 yaitu calculator chmod dengan ditambahkan awalan 0

	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createNewFile("sample.log", "this is sample log\nthis is sample log\nthis is sample log")
	addToFile("sample.log", "\nBang dodo ingin ke Bali")
	dataFile, err := readFile("sample.log")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dataFile)
}
