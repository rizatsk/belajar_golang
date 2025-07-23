package golang_gorutines

import (
	"fmt"
	"strconv"
	"testing"
)

func Goroutine_Range(channel chan<- string) {
	for i := 0; i < 10; i++ {
		channel <- "Perulangan ke " + strconv.Itoa(i)
	}
	close(channel)
}

func Test_Goroutine_Range(test *testing.T) {
	channel := make(chan string)

	go Goroutine_Range(channel)

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("Done")
}
