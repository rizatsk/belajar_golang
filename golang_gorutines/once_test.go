package golang_gorutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func CounterIncrement() {
	counter++
}

func OnlyOnce(wg *sync.WaitGroup, once *sync.Once) {
	defer wg.Done()
	once.Do(CounterIncrement)
}

func TestOnce(test *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go OnlyOnce(&wg, &once)
	}

	wg.Wait()
	fmt.Println(counter)
}
