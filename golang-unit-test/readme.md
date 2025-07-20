# Learner GO Unit test
Materi GO Unit test
https://docs.google.com/presentation/d/1XxMEaA-JsPHr9BUw2oIOPlEL_psI3EaUFUpuvdlDB_Q/mobilepresent?slide=id.p

## GO-LANG
### Rules Membuat Unit Test
- Untuk membuat unit test, nama file test harus berakhir `_test` contoh `hello_world_test.go`  
- Nama untuk nama function harus berawali `Test` contoh `TestHelloworld`  
- Harus memberikan paramter `*testing.T`  
- Tidak boleh memberikan return pada function unit test  

### Contoh pembuatan unit test
``` go
// helper/hello_world.go
package helper

func helloworld(name string) string {
	return "Hello " + name
}

// helper/hello_world_test.go
package helper

import "testing"

func TestHelloworld(test *testing.T) {
	result := helloworld("Rizat")
	if result != "Hello Rizat" {
		panic("Result is not Hello Rizat")
	}
}
```

### Menjalankan Unit Test
- Untuk menjalankan seluruh unit test semua yang ada `go test ./...`
- Untuk menjalankan seluruh unit test yang ada di dalam folder/package `go test`
- Lebih detail function bisa gunakan `go test -v`
- Jalankan unit test function tertentu bisa gunakan `go test -v -run TestNamaFunction`

### Meng-gagalkan unit test
Untuk meng-gagalkan unit test sebaiknya tidak menggunakan `panic` better menggunakan function error-nya yaitu `Fail()`, `FailNow()`, `Error()`, dan `Fatal()`.

#### test.Fail() and test.FailNow()
- Fail akan meng-gagalkan unit test, tapi tetap melanjutkan eksekusi unit test. maka hasil unit test dianggap gagal.
- FailNow akan meng-gagalkan unit test saat itu juga, tanpa melanjutkan eksekusi unit test nya.

#### test.Error(arguments) and test.Fatal(arguments)
- Error() - meng-gagalkan error dengan menampilkan log error, tetap menjalankan eksekusi unit test.
- Fatal() - meng-gagalkan error dengan menampilkan log error, tanpa menlanjutkan eksekusi unit test.

### Testify
Module untuk melakukan pengecekan unit test(Assertion), tidak lagi disarankan menggunakan if else
``` sh
go get github.com/stretchr/testify@v1.6.1
```

Contoh penggunaan-nya
``` go
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssertion(test *testing.T) {
	result := helloworld("Rizat")
	assert.Equal(test, "Hello Rizat", result)
}
```

#### Assert vs Require
- Assert ketika gagal akan tetap melanjutkan unit testnya.
- Require ketika gagal tidak melanjutkan unit testnya.

``` go
result := helloworld("Rizat")
require.Equal(test, "Hello Rizat", result)
```

### Skip Test
Untuk skip unit test nya
``` go
test.Skip("Log kenapa di skip")
```

### Before And After Test
Eksekusi code sebelum dan sesudah menjalankan test  
Dengan menggunakan function bernama `TestMain` dengan parameter `testing.M`
Go-Lang akan mengeksekusi function ini tiap kali menjalankan unit test di package, `hanya dieksekusi 1x/Go-Lang package`  
Jadi kita bisa mengatur Before dan After sesuai dengan yang kita mau.  
``` go
func TestsMain(m *testing.M) {
	fmt.Println("Sebelum unit Test")

	m.Run() // eksekusi semua unit test

	fmt.Println("Sesudah unit test")
}
```

### Sub Test
Membuat unit test didalam functuion unit test, untuk membuat sub test bisa menggunakan function `test.Run()`
``` go
func TestSubTest(test *testing.T) {
	t.Run("Rizat", func(t *testing.T) {
		result := HelloWorld("Rizat")
		assert.Equal(t, "Hello Rizat", result)
	})
	t.Run("Sakmir", func(t *testing.T) {
		result := HelloWorld("Sakmir")
		assert.Equal(t, "Hello Sakmir", result)
	})
}
```

### Table Test
Konsep test dengan menyediakan data berupa slice yang berisi parameter dan ekspektasi hasil dari unit test, menggunakan sub test.  
``` go
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
```

### Mock
Untuk mock function agar sesuai yang kita inginkan, contoh function yang disarankan untuk di mock yaitu call API, get data form Database, kalau function biasa itu tidak perlu untuk dimocking, seperti function validasi.  
- Bisa dilihat pada code `repository/category_repository_mock.go`
- dan untuk membuat function unit test di `service/category_service_test.go`


### Benchmark
Menghitung kecepatan performa kode aplikasi, dilakukan secara otomatis dengan iterasi kode sampai waktu tertentu dan sudah diatur oleh `testing.B` bawaan package dari `testing`, untuk membuat-nya dengan nama function diawali kata `Benchmark`, contoh `BanchmarkHelloworld`
``` go
func BenchmarkHelloworld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Helloworld("Rizat")
	}
}
```  
  
#### Running Benchmark
- `go test -v -bench=.` ini akan menjalankan semua unit test dan benchmark yang ada didalam module.
- `go test -v -run=NorMathUnitTest -bench=.` ini akan menjalankan hanya benchmark saja didalam module.
- `go test -v -run=NotMathUnitTest -beanch=BanckMarkNameTest` ini akan menjalankan benchmark yang ingin kalian jalankan.
- `go test -v -bench=./...` ini akan menjalankan benchamrk yang ada disemua module

#### Sub Benchmark
Dengan menggunakan function `Run()` sama seperti `sub test`
``` go
func BenchmarkHelloworld(b *testing.B) {
	b.Run("Rizat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Helloworld("Rizat")
		}
	})
	b.Run("Software Engineer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Helloworld("Software Engineer")
		}
	})
}
```

