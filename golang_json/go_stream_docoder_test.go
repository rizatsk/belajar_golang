package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(test *testing.T) {
	reader, _ := os.Open("user.json")
	decoder := json.NewDecoder(reader)

	var user map[string]interface{}
	_ = decoder.Decode(&user)

	fmt.Println(user["name"])
	fmt.Println(user["job"])
	fmt.Println(user["mobile_number"])
}
