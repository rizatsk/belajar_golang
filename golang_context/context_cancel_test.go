package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// Function Goroutine Leak
func CraeteCounter(context context.Context, wg *sync.WaitGroup) chan int {
	destination := make(chan int)

	wg.Add(1)
	go func() {
		defer func() {
			close(destination)
			wg.Done()
		}()

		counter := 1
		for {
			select {
			case <-context.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(test *testing.T) {
	var wg sync.WaitGroup
	parentContext := context.Background()
	context, cancel := context.WithCancel(parentContext)

	fmt.Println("Total goroutine =", runtime.NumGoroutine())

	destination := CraeteCounter(context, &wg)
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()

	wg.Wait()
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
}
