package golang_gorutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeTicker(test *testing.T) {
	fmt.Println("Run every 1 seconds")
	ticker := time.NewTicker(1 * time.Second)

	looping := 0
	for time := range ticker.C {
		fmt.Println(time)
		if looping == 5 {
			fmt.Println("Ticker di stop")
			ticker.Stop()
			break
		}
		looping += 1
	}
}

func TestTimeTick(test *testing.T) {
	fmt.Println("Run every 1 seconds")
	channel := time.Tick(1 * time.Second)

	looping := 0
	for time := range channel {
		fmt.Println(time)
		if looping == 3 {
			fmt.Println("Tick di stop")
			break
		}
		looping += 1
	}
}
