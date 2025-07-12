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

