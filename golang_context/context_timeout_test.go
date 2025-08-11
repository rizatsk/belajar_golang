package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Function Goroutine Leak
func CraeteCounterUseContextTimeout(context context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer func() {
			close(destination)
		}()

		counter := 1
		for {
			select {
			case <-context.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulation slow response
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(test *testing.T) {
	parentContext := context.Background()
	context, cancel := context.WithTimeout(parentContext, 5*time.Second)
	defer cancel()

	fmt.Println("Total goroutine =", runtime.NumGoroutine())

	destination := CraeteCounterUseContextTimeout(context)
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutine =", runtime.NumGoroutine())
}
