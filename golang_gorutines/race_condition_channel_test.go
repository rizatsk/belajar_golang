package golang_gorutines

import (
	"fmt"
	"testing"
)

func IncrementChannelX(channel chan<- int) {
	sum := 0
	for j := 1; j <= 100; j++ {
		sum++
	}

	channel <- sum
}

func TestRaceConditionChannel(test *testing.T) {
	total_goroutine := 1000
	channel := make(chan int, total_goroutine)

	for i := 1; i <= total_goroutine; i++ {
		go IncrementChannelX(channel)
	}

	total := 0
	for i := 1; i <= total_goroutine; i++ {
		total += <-channel
	}

	fmt.Println("Counter =", total)
}
