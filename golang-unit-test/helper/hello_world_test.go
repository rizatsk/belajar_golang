package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloworld(test *testing.T) {
	result := helloworld("Rizats")
	if result != "Hello Rizat" {
		test.Error("Result must be Hello Rizat")
	}
}

func TestHelloWorldAssertion(test *testing.T) {
	result := helloworld("Rizats")
	assert.Equal(test, "Hello Rizat", result)
}

func TestTableTest(test *testing.T) {
	usersTest := []struct {
		name, request, expected string
	}{
		{
			name:     "helloworld(Rizat)",
			request:  "Rizat",
			expected: "Hello Rizat",
		},
		{
			name:     "helloworld(Sakmir)",
			request:  "Sakmir",
			expected: "Hello Sakmir",
		},
	}

	for _, user := range usersTest {
		test.Run(user.name, func(test *testing.T) {
			result := helloworld(user.request)
			assert.Equal(test, user.expected, result)
		})
	}
}
