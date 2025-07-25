package golang_gorutines

import (
	"fmt"
	"testing"
	"time"
)

func Helloworld() {
	fmt.Println("Hello world")
}

func TestHelloworld(test *testing.T) {
	go Helloworld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutines(test *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
