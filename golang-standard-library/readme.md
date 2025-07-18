# Learner GO Standard Libarary
Materi GO
https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?usp=sharing

## GO-LANG
### Standard Library
Package-package yang sudah disediakan oleh Go-lang itu sendiri  
🚀 https://pkg.do.dev/std

### Fmt(Format)
``` go
firstName := "Rizat"
lastName := "Sakmir"

fmt.Printf("Hello '%s %s'\n", firstName, lastName)
```

### Errors
``` go
var (
	validationError = errors.New("Is validation error")
	notFoundError   = errors.New("Is not found error")
)

func cekName(name string) error {
	if name == "" {
		return validationError
	} else if name != "Agus" {
		return notFoundError
	}

	return nil
}

func main() {
	errorCekName := cekName("Rizat")
	if errorCekName != nil {
		if errors.Is(errorCekName, validationError) {
			fmt.Println("Validation error")
		} else if errors.Is(errorCekName, notFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unkown error")
		}
	}
}
```

### OS
``` go
import {
	"os"
}
```

### Flag
Untuk membuat flag  
``` go
import (
	"flag"
	"fmt"
)

func main() {
	// (key, default value, description)
	username := flag.String("username", "root", "Database username")
	password := flag.String("password", "root", "Database password")

	flag.Parse()

	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
}
```
Cara menjalankan dengan flag, example:
`go run flag.go -username=localhost -password=rahasia`

### String
Untuk manipulasi tipe data string
``` go
import {
	"strings"
}
```

### Strconv
String convertion untuk merubah tipe data integer ke string, begitu sebaliknya
``` go
import {
	"strconv"
}
```

### Math
Package berisikan fungsi matematika
``` go
import {
	"math"
}
```

### Container / List
Package container/list adalah implementasi struktur data double linked list di Go-Lang
``` go
import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Rizat")
	data.PushBack("Sakmir")
	data.PushBack("Software Engineer")

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
```

### Contaienr / Ring
Package untuk implementasi struktur data circular list, dimana element awal akan terhapus, dan masuk element baru
``` go
import (
	"container/ring"
	"fmt"
	"time"
)

func main() {
	data := ring.New(5)
	aktivitas := []string{
		"Membuka aplikasi",
		"Melihat profil pengguna",
		"Mengirim pesan",
		"Menerima notifikasi",
		"Melakukan pembayaran",
		"Mencari produk baru", // Ini akan menggantikan aktivitas pertama
	}

	fmt.Println("-- Menambahkan atktifiatas ke dalam ring ---")
	for _, akt := range aktivitas {
		now := time.Now().Format("2006-01-02 15:04:05")
		data.Value = fmt.Sprintf("[%s] %s", now, akt)

		fmt.Println(data.Value)

		data = data.Next()
	}

	fmt.Println("-- Riwayat 5 aktifitas terakhir ---")
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
```

### Sort
Sort digunakan untuk proses pengurutan, dimana implementasinya harus dibuat kontrak interface dahulu `Len, Less, Swap` lalu tinggal di sort dengan package dari `sort.Sort`
``` go
import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (slice UserSlice) Len() int {
	return len(slice)
}

func (slice UserSlice) Less(i, j int) bool {
	return slice[i].Age < slice[j].Age
}

func (slice UserSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	users := []User{
		{Name: "Rizat", Age: 25},
		{Name: "Joko", Age: 35},
		{Name: "Bambamg", Age: 30},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
```

### Time
Package untuk waktu, untuk format tanggal `yyyy-mm-dd hh:mm:ss` itu harus menggunakan format `2006-01-02 15:04:05`
``` go
import {
	"time"
	"fmt"
}

func main() {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	acitvity = fmt.Sprintf("[%s] %s", now, "Sedang belajar")

	timeString = "2025-07-12 18:58:05"
	timeParse = time.Parse("2006-01-02 15:04:05", timeString)
	fmt.Println(activity)
}
```

### Reflect
Untuk melihat struktur code saat aplikasi sedang berjalan, bisa dibuat untuk validasi
``` go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name, Email, Phone string `required:"true"`
}

func isValid(value any) (result bool) {
	result = true
	dataType := reflect.TypeOf(value)
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	isValidUser := isValid(User{
		Name:  "Dodo",
		Email: "dodo@gmail.com",
		Phone: "081231412314",
	})
	fmt.Println("Validasi user :", isValidUser)
}
```

### Regexp
``` go
import {
	"regexp"
}
```

### Encoding
``` go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Rizat Sakmir"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		stringDecoded := string(decoded)
		fmt.Println(stringDecoded)
	}
}
```

### Slices
``` go
import {
	"slices"
}

func main() {
	names := []string{"Dodo", "Rizat", "Hengki"}
	values := []string{"100", "90", "60"}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Dodo"))
	fmt.Println(slices.Index(names, "Rizat"))
}
```

### Path
Untuk manipulasi data path di URL/Path file system
``` go
import {
	"path"
	"fmt"
}

func main() {
	fmt.Println(path.Dir("hello/word.go"))
	fmt.Println(path.Base("hello/word.go"))
	fmt.Println(path.Ext("hello/word.go"))
	fmt.Println(path.Join("dodo", "hengki", "main.go"))
}
```

Khusus untuk manipulasi path file system di windows menggunakan `path/filepath`
``` go
import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/word.go"))
	fmt.Println(filepath.Base("hello/word.go"))
	fmt.Println(filepath.Ext("hello/word.go"))
	fmt.Println(filepath.Join("dodo", "hengki", "main.go"))
}
```

### BUFIO
Untuk input output di Golang, seperti Rider dan Writer  
Reader
``` go
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Reader on file
	input := strings.NewReader("This is long string\nwith new line \n")
	reader := bufio.NewReader(input) // Tujuan read

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	// Writer
	writer := bufio.NewWriter(os.Stdout) // Tujuan write
	writer.WriteString("Hello world\n")
	writer.WriteString("Selamat belajar golang\n")
	writer.Flush()
}
```

### File Manipulation
Membuat file, baca file, dan menambahkan isi file  
Untuk calculate permission bisa menggunakan https://chmod-calculator.com  
``` go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666) // 0666 yaitu calculator chmod dengan ditambahkan awalan 0

	if err != nil {
		return err
	}

	defer file.Close() // di close agar tidak terjadi memory leak
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createNewFile("sample.log", "this is sample log\nthis is sample log\nthis is sample log")
	addToFile("sample.log", "\nBang dodo ingin ke Bali")
	dataFile, err := readFile("sample.log")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dataFile)
}

```
