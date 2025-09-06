package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ResponseAPITag struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func TestJSONTag(test *testing.T) {
	logJSON(ResponseAPITag{
		Status:  "success",
		Message: "Success test used struct",
		Data: []User{
			{
				Name: "Rizat Struct",
				Job:  "Software Engineer",
			}, {
				Name: "Hengki Struct",
				Job:  "Admin Panel",
			},
		},
	})

	json_tag := `{"status":"success","message":"Success test used struct","data":[{"Name":"Rizat Struct","Job":"Software Engineer"},{"Name":"Hengki Struct","Job":"Admin Panel"}]}`
	jsonBytes := []byte(json_tag)

	response := &ResponseAPITag{}
	json.Unmarshal(jsonBytes, response)

	fmt.Println(response)
	fmt.Println(response.Status)
}
