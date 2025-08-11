package golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateContext(test *testing.T) {
	background := context.Background()
	fmt.Println("Value is", background)

	todo := context.TODO()
	fmt.Println("Value is", todo)
}
