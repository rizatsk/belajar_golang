package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(test *testing.T) {
	writer, _ := os.Create("sample_stream_encoder.json")
	encoder := json.NewEncoder(writer)

	user := User{
		Name: "Rizat",
		Job:  "Software Engineer",
	}

	_ = encoder.Encode(user)

	fmt.Println(user.Name)
	fmt.Println(user.Job)
}
