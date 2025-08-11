package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Function Goroutine Leak
func CraeteCounterUseContextDeadline(context context.Context) chan int {
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

func TestContextWithDeadline(test *testing.T) {
	parentContext := context.Background()
	context, cancel := context.WithDeadline(parentContext, time.Now().Add(5*time.Second))
	defer cancel()

	fmt.Println("Total goroutine =", runtime.NumGoroutine())

	destination := CraeteCounterUseContextDeadline(context)
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutine =", runtime.NumGoroutine())
}
