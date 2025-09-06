package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshal(test *testing.T) {
	jsonRequest := `{"Name":"Hengki Struct","Job":"Admin Panel"}`
	jsonBytes := []byte(jsonRequest)

	user := &User{}
	json.Unmarshal(jsonBytes, user)

	fmt.Println(user.Name)
	fmt.Println(user.Job)
}
