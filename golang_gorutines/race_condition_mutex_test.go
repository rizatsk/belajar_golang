package golang_gorutines

import (
	"fmt"
	"sync"
	"testing"
)

func IncrementMutexX(x *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1; j <= 100; j++ {
		mu.Lock()
		*x += 1
		mu.Unlock()
	}
}

func TestRaceConditionMutex(test *testing.T) {
	x := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go IncrementMutexX(&x, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("Counter =", x)
}
