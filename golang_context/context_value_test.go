package golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(test *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Get value
	fmt.Println("\nGet value in context")
	fmt.Println(contextA.Value("b")) // Tidak dapat ambil dari child
	fmt.Println(contextD.Value("b")) // Dapat ambil dari parent
	fmt.Println(contextF.Value("f")) // Dapat
}
