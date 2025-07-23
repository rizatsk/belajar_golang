package golang_gorutines

import (
	"fmt"
	"testing"
)

func SendChannel(channel chan<- string) {
	fmt.Println("Hello is run channel")
	channel <- "Rizat" // kirimkan data ke channel
	fmt.Println("Success send data channel")
}

func TestCreateChannel(test *testing.T) {
	channel := make(chan string)
	defer close(channel) // menutup channel agar tidak terjadi memory leak

	go SendChannel(channel)
	fmt.Println("Run after goroutine")

	name := <-channel // menerima data channel
	fmt.Println("Your name", name)
}
