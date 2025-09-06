# Learner GO JSON
Materi GO Context
https://docs.google.com/presentation/d/1mGVdO7Khmiw-9lDmkCWqd9l5_UqohKRkpHbjabRD--U

## GO-LANG
### Pengenalan
Go (Golang), Belajar Golang Jason

### Encode
Untuk conversi data ke JSON
``` go
bytes, err := json.Marshal(data)
if err != nil {
    panic(err)
}
fmt.Println(string(bytes))
```

### Decode
Untuk conversi data dari JSON ke Type data Go
``` go
jsonRequest := `{"Name":"Hengki Struct","Job":"Admin Panel"}`
jsonBytes := []byte(jsonRequest)

user := &User{}
json.Unmarshal(jsonBytes, user)

fmt.Println(user.Name)
fmt.Println(user.Job)
```

### Json Tag
Standar penulisan map key di Go menggukanan PascalCase,  
sedangkan standar dari JSON itu menggunakan snake_case/cammelCase  
membuat keliruan saat integrasi, maka dari itu kita bisa membuat json Tag dimana akan diconversi otomatis saat melakukan encode dan decode.
``` go
type ResponseAPITag struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
```

### Stream Decode
Tidak perlu masukan data json ke dalam string alias membuat variable baru, langsung bisa di decode jadi tipe data Go cocok untuk baca data file JSON, request dari API, dll
``` go
func TestStreamDecoder(test *testing.T) {
	reader, _ := os.Open("user.json")
	decoder := json.NewDecoder(reader)

	var user map[string]interface{}
	_ = decoder.Decode(&user)

	fmt.Println(user["name"])
	fmt.Println(user["job"])
	fmt.Println(user["mobile_number"])
}
```

### Stream Encoder
Tidak perlu lagi membuat variable dan string kan lagi hasil data Encodernya, cocok untuk membuat file json ke dalam file atau response API
``` go
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
```