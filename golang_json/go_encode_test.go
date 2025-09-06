package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type ResponseAPI struct {
	Status  string
	Message string
	Data    interface{}
}

type User struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func TestMarshal(test *testing.T) {
	logJSON("Rizat")
	logJSON(1)
	logJSON(true)
	logJSON(map[string]interface{}{
		"status":  "success",
		"message": "Success test marshel",
		"data": []map[string]interface{}{
			{
				"name": "Rizat Sakmir",
				"job":  "Software Engineer",
			}, {
				"name": "Hengki",
				"job":  "Admin Panel",
			},
		},
	})
	logJSON(ResponseAPI{
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
	logJSON([]string{"Rizat", "Sakmir", "Software Engineer"})
}
