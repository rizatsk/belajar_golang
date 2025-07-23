package golang_gorutines

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Rizat"
	channel <- "Software"
	channel <- "Engineer"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Done")
}
