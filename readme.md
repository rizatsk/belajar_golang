# Learner GO Dasar
Materi GO
https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?usp=sharing

Discord group
https://discord.gg/pEgSFrpPSv

## Install GO
Install GO di website resmi https://go.dev/

## Learner GO Dasar
### Cek instalasi GO
``` go
go version
```

### Create module in Golang
``` go
go mod init nama-module
```

### Running GO Development
``` go
go run name-file-go.go
```

### Running GO Production
``` go
go build
```
Akan menghasilkan file dengan extention yang bisa dijalankan dari setiap OS

### Tipe Data
Otomatis akan di convert ke tipe data yang sesuai dalam penggunaan code
#### Number
1. Integer
- int8 = -128 => 255, dst
- uint8 = 0 => 255, dst
2. Floating point
- float32
- float64
- complex64
- complex128
3. Alias
- byte = uint8
- rune = int32
- int = minimal int32
- uint = minimal uint32

#### Boolean
- true 
- false

#### String
Dengan diawali dengan `"` diakhiri `"` example: `"Hello world"`
``` go
// Hitung jumlah string
len("string")

// Ambil karakter di posisi string dimulai dari 0
"string"[number] // Ini akan menjadi byte
```

### Variable
Tempat data untuk menampung data
``` go
var name string

name = "Hello world"
fmt.Println(name); // Print "Hello world"

name = "Hello Rizat"
fmt.Println(name); // Print "Hello Rizat"
```

Otomatis dibuatkan type data variable
``` go
var name = "Rizat"

// Bisa juga seperti ini untuk declarasi variable
job_user := "Software Engineer"
```

Multiple declarasi variable
``` go
var (
  firstName = "Rizat"
  lastName = "Sakmir"
)
```

### Contant
Variable yang tidak bisa lagi diubah datanya.
``` go
const database = "MySQL"
```

Multiple deklarasi const
``` go
const {
  firstName = "Rizat"
  lastName = "Sakmir"
}
```

### Conversion Data
``` go
var nilai32 = 32768
var nilai64 = int64(nilai32)

const convertString = "Rizat Sakmir"
var s = string(convertString[6])
```

### Type Declaration
``` go
type NoKTP string

var ktpRz NoKTP = "0001111";
var ktpExample NoKTP = NoKTP(ktpRz)
```

### Operation Matematika
`+ , - , * , / , %`

Augmented Assignmenst
`a = a + 10 => a += 10`

### Unerary Operator
`++ , -- , - , + , !`

### Operator Perbandingan
`> , < , >= , <= , == , !=`

### Operator Boolean
`&& , || , !`

### Array
Isi data array tidak bisa bertambah setelah dibuat, kalau ingin flexible bisa gunakan `slice`.
dan hanya bisa 1 tipe data yang sama.
example:
``` go
var numbers = [3]int{1, 2, 3}
var name_users = [2]string{"Rizat", "Sakmir"}
fmt.Println(name_users[1])
```

- Kalau ingin type data yang berbeda
``` go
var mixed [3]interface{}
mixed[0] = 42
mixed[1] = "hello"
mixed[2] = true

fmt.Println(mixed)
```

### Slice
Potongan data array yang flexible ukuranya bisa berubah, atau dia bisa buat array flexible
`notes*: Ketika membuat slice baru dari array dan rubah data slice nya, data array nya juga ikut berubah`
``` go
array := [...]int{1, 2, 3}
slice := []int{1, 2, 3}

var name_users = [3]string{"Rizat", "Sakmir", "Joko"}
var slice_name_users_0_1 = name_users[0:1]
var slice_name_users_1_sdt = name_users[1:]
var slice_name_users_2_sampai_ke_depan = name_users[:2]
var slice_name_users_0_sdt = name_users[:]

// Membuat slice dari awal
newSlice := make([]string, 2, 3) // (typeData, panjang, kapasitas)
newSlice[0] = "Rizat"
newSlice[1] = "Rizat"

fmt.Println(newSlice)
fmt.Println(len(newSlice))
fmt.Println(cap(newSlice))
```

#### Function slice
- len(slice) - Untuk mendapatkan panjang
- cap(slice) - Untuk mendapatkan kapasitas
- append(slice, data) - Membuat slice baru dengan menambahkan data ke posisis terakhir, jika kapasitas sudah penuh, maka akan membuat array baru
- make([]TypeData, length, capacity) - Membuat slice baru
- copy(destination, source) - Menyalin slice dari source ke destination


### Map
Bentuk data object terdapat key dan value, kalau tidak ada key yang diambil akan menggunakan defautl value tergantung dengan type data yang awal di deklrasi
``` go
person := map[string]string {
  "name": "Rizat",
  "address": "Jakarta"
}

fmt.Println(person["name"])
```

#### Function Map
- len(map) - Untuk mendapatkan jumlah data map
- map[key] - Ambil data map dari key
- map[key] = value - Ubah data map sesuai dengan key
- make(map[TypeKey]TypeValue) - Membuat map baru
- delete(map, key) - Delete data berdasarkan key dari map

### Perkondisian IF
Digunakan untuk perkondisian
``` go
name := "Rizat"

if name == "Rizat" {
  fmt.Println("Hallo Rizat")
} else if name == "Joko" {
  fmt.Println("Hello Joko")
} else {
  fmt.Println("Hello orang asing")
}

// If Short Statement
if length := len(name); length > 5 {
  fmt.Println("Nama terlalu panjang")
} else {
  fmt.Println("Nama sudah sesuai")
}
```

### Switch
``` go
name := "Rizat"

switch name  {
  case "Rizat":
    fmt.Println("Hallo Rizat")
  case "Jaka":
    fmt.Println("Hallo Jaka")
  default:
    fmt.Println("Hi orang asing")
}
```

### For Loops
``` go
counter := 1

for counter <= 10 {
  fmt.Println("Perulangan ke ", counter);
  counter++
}

for counter := 1; counter <= 10; counter++ {
  fmt.Println("Perulangan ke ", counter);
}

names := []string{"Bambang", "Joko", "Rizat"}
for index, name := range names{
  fmt.Println(name)
}
```

### Break And Continue Perulangan
Break digunakan untuk memberhentikan seluruh perulangan.  
Continue digunakan untuk memberhitkan perulangan, dan melanjutkan perulangan berikutnya.
``` go
for counter := 1; counter <= 10; counter++ {
  if counter == 6 {
    break
  }

  if counter == 2 {
    continue
  }

  fmt.Println("Perulangan ke ", counter)
}
// Result 1, 3, 4, 5
```

### Function

``` go
func name_function() {
  fmt.Println("This is name_function")
}

func main() {
  name_function()
  fmt.Println("Running program Golang")
}
```
``` go
project/
│
├── main.go
└── slice.go
```
Untuk berbeda file hanya cukup panggil dengan nama functionya saja.  
Otomatis akan bisa digunakan, dan cara untuk menjalankan mode developer cukup `docker run .`  
Jadi hanya menjalankan function main.

#### Berbeda hirarki file
``` go
project/
│
├── main.go
└── utils/
    └── slice.go
```

``` go
// utils/slice.go
package utils

import "fmt"

func PrintSlice() {
    fmt.Println("Ini dari utils/slice.go")
}

// main.go
package main

import (
    "your_module_name/utils" // module name sesuai dengan yang ada di go.mod
)

func main() {
    utils.PrintSlice()
}
```

#### Function Parameter
``` go
// Function parameter
func sayHello(name string, job string) {
  fmt.Println("Hello ", name);
  fmt.Println("Your jobs is ", job);
}

// Function parameter dengan object
type User struct {
	ID   int
	Name string
}

func sayHelloObject(user User) {
	fmt.Println("Running function say hello object")
	fmt.Println("User id", user.ID)
	fmt.Println("Your name", user.Name)
}

func main() {
  sayHello("Rizat", "Software Engineer")
  sayHelloObject(User{ID: 1, Name: "Rizat"})
}
```

### Function return value
``` go
func getHello(name string) string {
  return "Hello" + name
}

func multipleValues() (string, string) {
  return "Multiple", "Hello"
}

func main() {
  result := getHello("Rizat")
  fmt.Println(result)

  multi, say := multipleValues()
  fmt.Println(multi, say)
}
```

### Parameter Function Variadic
Dimana parameter function di masukan sebanyakan apapun dengan cara :
``` go
func sumAll(numbers ...int) int {
  total := 0
  for _, number := range numbers {
    total += number
  }

  return total
}

func main() {
  total := sumAll(10, 10, 10, 10);
  fmt.Println(total)
}
```

#### Parameter Function Slice
``` go
numbers := []int{1, 2, 3, 4};
fmt.Println(sumAll(numbers...))
```

### Function Di Jadikan Variable
``` go
func getGoodBye(name string) string {
  return "Good bye" + name
}

func main() {
  goodbye := getGoodBye
  fmt.Println(goodbye("Rizat"))
}
```

### Function As Parameter
``` go
func sayHelloWithFilter(name string, filter func(string) string) {
  fmt.Println("Hello", filter(name));
}

func spamFilter(name string) string {
  if name == "anjing" {
    return "..."
  } else {
    return name
  }
}

func main() {
  sayHelloWithFilter("Joko anjing", spamFilter)
}
```

#### Function dengan type declaration
``` go
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
  fmt.Println("Hello", filter(name));
}
```


### Defer
Defer function untuk menjadwalkan function dieksekusi terakhir seperti halnya finally di javascript
``` go
func logging() {
  fmt.Println("Selesai di eksekusi")
}

func main() {
  defer logging()
  fmt.Println("Run application")
}
```

### Panic
Panic function untuk menghentikan program, tapi defer akan tetap dieksekusi
``` go
panic("ERROR")
```

### Recover
Recover function untuk menangkap error panic, sehingga program tetap berjalan
``` go
func endApp() {
  fmt.Println("End App")
  message := recover()
  fmt.Println("Terjadi error", message)
}

func main() {
  defer endApp()
  panic("ERROR")
}
```

### Struct
Yaitu type data object  
Struct mendefinisikan bentuk data dan dapat memiliki metode yang melekat padanya.
``` go
type Customer struct {
  Name, Adresss string
  Age int
}

func main() {
  var dodo Customer
  dodo.Name = "Dodo"
  dodo.Adress = "Indonesia"
  dodo.Age = 30

  fmt.Println(dodo)
}
```

#### Struct Method
Struct method ini agar function menempel pada type struct
``` go
type Customer struct {
  Name, Adresss string
  Age int
}

func (customer Customer) sayHello(name string) {
  fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
  var dodo Customer
  dodo.Name = "Dodo"
  dodo.Adress = "Indonesia"
  dodo.Age = 30

  fmt.Println(dodo)

  customer.sayHello("Hengki")
}
```

### Interface
Tipe data abstrak
``` go
type HashName interface {
  GetName() string
}

type Person struct {
  Name string
}

func sayHello(value HashName) {
  fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string {
  return person.Name
}

func main() {
  person := Person(Name: "Bobby")
  sayHello(person)
}
```
Untuk tipe data yang tidak memiliki tipe data yang pasti seperti panic, dan recover bisa menggunakan `any` / `interface{}`

### Tipe Data Nil
Pada golang tipe data maka secara otomatis dibuatkan default valuenya, sesuai dengan tipe data yang diberikan.
`Nil yaitu data kosong`, tapi hanya bisa digunakan oleh beberapa tipe data saja seperti `interface, function, map, slice, pointer dan channel`.
``` go
func newMap(name string) map[string] string {
  if name == "" {
    return nil
  } else {
    return map[string] string {
      "name": name
    }
  }
}
```

### Type Assertions
Untuk merubah tipe data menjadi tipe data yang diingkan.  
biasa digunakan untuk interface type any
``` go
func random() interface{} {
  return "OK"
}

func main() {
  result := random()
  resultString := result.(string)
  fmt.Println(resultString)
}
```

#### Type Assestions Dengan Switch
``` go
func main() {
  result := random()
  switch value := result.(type) {
    case string:
      fmt.Prinln("String", value)
    case int:
      fmt.Println("Int", value)
    default:
      fmt.Println("Unknown")
  }
}
```

### Pointer
Di GO ketika membuat variable baru pass by variable lama itu pass by value, dimana variable yang dicopy tidak berpengaruh ketika diubah datanya.  
Fitur untuk pass by reference jadi menggunakan memory variable yang sama. menggunakan `&` disebelum variablenya example: `&address1`
``` go
func main() {
  address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
  address2 := &address1

  address2.City = "Bandung"

  fmt.Println(address1) // {"Bandung", "Jawa Barat", "Indonesia"}
  fmt.Println(address2) // {"Bandung", "Jawa Barat", "Indonesia"}
}
```

### Pointer Di Function
Parameter di function ditambahkan `&`
``` go
func ChangeAddressToIndonesia(address Address) {
  address.Country = "Indonesia"
}

func main() {
  address := Address{"Subang", "Jawa Barat", ""}
  ChangeAddressToIndonesia(&address)

  fmt.Println(address) // {"Subang", "Jawa Barat", "Indonesia"}
}
```

### Pointer Di Method
Untuk dapat merubah data Man di method bisa tambahkan `*` pada type-nya.
``` go
type Man struct {
  Name String
}

func (man *Man) Married() {
  man.Name = "Mr." + man.Name
}

func main() {
  rizat := Man("Rizat")
  rizat.Married()

  fmt.Println(rizat)
}
```

### Asterisk Operator Atau *
Operator `*` penggunaan dengan menambahkan di awal variable dengan `*variable`
Agar variable yang awalnya ber pointer ke variable awal, variable awal juga ikut berubah, example :
``` go
func main() {
  address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
  address2 := &address1

  address2.City = "Bandung"
  *address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

  fmt.Println(address1) // {"Jakarta", "DKI Jakarta", "Indonesia"}
  fmt.Println(address2) // {"Jakarta", "DKI Jakarta", "Indonesia"}
}
```

### Operator New
Untuk membuat pointer baru sama saja seperti penggunaan `&` dengan penggunaan `new`
``` go
func main() {
  address1 := new(Address)
  address2 := address1

  address2.Country = "Bandung"

  fmt.Println(address1) // &{ "Indonesia"}
  fmt.Println(address2) // &{ "Indonesia"}
}
```

### Package
Folder untuk menemaptkan code di direktori folder  
Nama package mengikuti nama foldernya
``` go
project/
│
└── main.go
└── utils/
    └── slice.go

// slice.go
pakcage utils

func Slice() {
  return "This is slice"
}

// main.go
import (
    "your_module_name/utils" // module name sesuai dengan yang ada di go.mod
    "fmt"
)

func main() {
  fmt.Println(utils.Slice())
}
```

### Access Modifier
Yaitu untuk dapat diakses dari package lain penamaan menggunakan awalan huruf besar `func Example() {}` kalau tidak ingin bisa diakses gunakan huruf kecil `func example() {}`.  
notes*: Tidak hanya function bisa juga untuk variable, struct, interface dll.  

### Package Initialization
Yaitu untuk meninitialisasi package sebelum function di gunakan
``` go
var connection string

func init() {
  connection = "MYSQL"
}

func GetDatabase() string {
  return connection // akan berisi MYSQL
}
```

#### Blank Package Initialization
Ketika hanya import package yang hanya Initialization-nya saja saat di import bisa menggunakan `_`
``` go
import {
  fmt
  _"your_module_name/blank_package"
}
```

### Error
Menggunakan package `errors` yang sudah disediakan oleh GO
```go
import {
  "errors"
}

func Pembagian(nilai int, pembagi int) (int, error) {
  if pembagi == 0 {
    return 0, errors.New("Pembagian Dengan NOL")
  } else {
    return nilai / pembagi, nil
  }
}

func main() {
  hasil, err := Pembagian(100, 10)
  if err == nil {
    fmt.Println("Hasil", hasil)
  } else {
    fmt.Println("Error", err.Error())
  }
}
```

### Rules in GO
#### Tidak boleh menggunakan nama funtion yang sama walaupun di file yang berbeda: example
``` go
helloworld.go
  func main() {}

example.go
  func main() {}
```

#### Jika variable yang tidak digunakan bisa menggunakan _
``` go
func multipleValues() (string, string) {
  return "Multiple", "Hello"
}

func main() {
  multi, _ := multipleValues()
  fmt.Println(multi)
}
```

