# Learner GO Context
Materi GO Context
https://docs.google.com/presentation/d/1WhJvRpKPWq7LY9P6fMN93vkjKa1bJwBQebbieKdefPw/edit?slide=id.p#slide=id.p

## GO-LANG
### Pengenalan
Go (Golang), context adalah sebuah objek yang digunakan untuk meneruskan informasi dan kendali antar goroutine atau fungsi, terutama dalam konteks operasi yang memerlukan pembatalan (cancellation), batas waktu (timeout), dan nilai-nilai terkait permintaan (request-scoped values)

### Interface Context
Untuk membuat context itu menggunakan interface, contoh:
``` go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interaface{}) interface{}  
}
```

### Context Background
`context.Background()` Membuat context kosong, tak pernah dibatalkan, tak pernah time out, dan tidak memiliki value. Biasanya digunakan main function, dalam unit test, dalam awal proses request.

### Context TODO
`context.TODO()` Membuat context kosong, biasa digunakan ketika belum jelas context apa yang ingin dilakukan.

### Parent and Child Context
Seperti bahasa pemograman OOP 1 parent bisa memiliki banyak child, dan child hanya bisa memiliki 1 parent.

### Contect With Value
Dengan menggunakan Pair ada key dan value, saat menambahkan value ke context secara otomatis akan membuat child context baru, original textnya tidak berubah, untuk membuatnya dengan `contect.WithValue(parent, key, value)`
``` go
func TestContextWithValue(test *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Get value
	fmt.Println("\nGet value in context")
	fmt.Println(contextA.Value("b")) // Tidak dapat ambil dari child
	fmt.Println(contextD.Value("b")) // Dapat ambil dari parent
	fmt.Println(contextF.Value("f")) // Dapat
}
```

### Context With Cancel 
Untuk membatalkan proses yang berjalan untuk membuatnya menggunakan `context.WithCancel(parent)`  
``` go
// Function Goroutine Leak
func CraeteCounter(context context.Context, wg *sync.WaitGroup) chan int {
	destination := make(chan int)

	wg.Add(1)
	go func() {
		defer func() {
			close(destination)
			wg.Done()
		}()

		counter := 1
		for {
			select {
			case <-context.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(test *testing.T) {
	var wg sync.WaitGroup
	parentContext := context.Background()
	context, cancel := context.WithCancel(parentContext)

	fmt.Println("Total goroutine =", runtime.NumGoroutine())

	destination := CraeteCounter(context, &wg)
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()

	wg.Wait()
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
}
```

### Context With Timeout
Untuk sinyal cancel secara otomatis dengan timeout, biasa digunakan untuk query ke database atau http API. Dengan menggunakan `context.WithTimeout(parent, duration)`
``` go
// Function Goroutine Leak
func CraeteCounterUseContextTimeout(context context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer func() {
			close(destination)
		}()

		counter := 1
		for {
			select {
			case <-context.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulation slow response
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(test *testing.T) {
	parentContext := context.Background()
	context, cancel := context.WithTimeout(parentContext, 5*time.Second)
	defer cancel()

	fmt.Println("Total goroutine =", runtime.NumGoroutine())

	destination := CraeteCounterUseContextTimeout(context)
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutine =", runtime.NumGoroutine())
}
```

### Context With Deadline
Untuk memberikan timeout pada waktu yang ingin ditentukan bukan dari timeout waktu sekarang. misal timeout berjalan ketika waktu 10:00 PM. Untuk membuat nya menggunakan function `context.WithDeadline(parent, time)`
``` go
// Function Goroutine Leak
func CraeteCounterUseContextDeadline(context context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer func() {
			close(destination)
		}()

		counter := 1
		for {
			select {
			case <-context.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulation slow response
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(test *testing.T) {
	parentContext := context.Background()
	context, cancel := context.WithDeadline(parentContext, time.Now().Add(5*time.Second))
	defer cancel()

	fmt.Println("Total goroutine =", runtime.NumGoroutine())

	destination := CraeteCounterUseContextDeadline(context)
	fmt.Println("Total goroutine =", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutine =", runtime.NumGoroutine())
}

```