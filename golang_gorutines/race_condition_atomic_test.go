package golang_gorutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func IncrementX(x *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1; j <= 100; j++ {
		atomic.AddInt64(x, 1)
	}
}

func TestRaceCondition(test *testing.T) {
	var x int64 = 0
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go IncrementX(&x, &wg)
	}

	wg.Wait()
	fmt.Println("Counter =", x)
}
