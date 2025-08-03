package golang_gorutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(test *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Time now :", time.Now())

	time := <-timer.C
	fmt.Println("This is time :", time)
}

func TestAfter(test *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Time now :", time.Now())

	time := <-channel
	fmt.Println("This is time :", time)
}

func TestAfterFunc(test *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Execute after 5 second")
		wg.Done()
	})

	wg.Wait()
}
