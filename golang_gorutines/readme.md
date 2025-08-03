# Learner GO Gorutines
Materi GO Gorutines
https://docs.google.com/presentation/d/1A78dn_g6HfxfRor9XBUAGPQM9vT6_SnrQGrQ2z0myOo/edit

## GO-LANG
### Pengenalan
- Goroutine dijalankan scheduler di dalam thread, dia berjalan secara concurency(yaitu jalan dalam 1 thread tapi bisa bergantian proses-nya tidak perlu menunggu 1 proses selesai) dan ini sudah diatur otomatis oleh Goroutine tanpa harus kita config manual.  
- Jumlah thread nya biasanya sebanyak core di CPU.

### Membuat Goroutine
Cukup dengan tambahkan perintah `go` sebelum memanggil function yang akan dijalankan, dan akan dijalankan secara asynchronous, tidak ditunggu sampai selesai, dan akan menjalankan code selanjutnya.
``` go
func Helloworld() {
	fmt.Println("Hello world")
}

func TestHelloworld(test *testing.T) {
	go Helloworld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
```

### Channel
- Channel tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine.
- Terdapat pengirim dan penerima, biasanya pengirim dan penerima goroutine yang berbeda, saat mengirimkan data ke channel goroutine akan terblock sampai ada yang menerimanya.
- Channel cocok sekali untuk alternatif mekanis async await seperti dibahasa pemograman lain.

#### Membuat Channel
``` go
func SendChannel(channel chan string) {
	fmt.Println("Hello is run channel")
	channel <- "Rizat" // kirimkan data ke channel
	fmt.Println("Success send data channel")
}

func TestCreateChannel(test *testing.T) {
	channel := make(chan string)
	defer close(channel) // menutup channel agar tidak terjadi memory leak

	go SendChannel(channel)
	fmt.Println("Run after goroutine")

	name := <-channel // menerima data channel
	fmt.Println("Your name", name)
}
```

### Channel In and Out
Parameter channel di function bisa kita tandai channel digunakan untuk `in`(mengirim data) atau `out`(menerima data)
``` go
func OnlyIn(channel chan<- string) {
	channel <- "Rizat Sakmir"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
}
```

### Buffered Channel
Untuk menampung data antrian di Channel, `Buffer Capacity` kita bebas memasukan berapa jumlah kapasitas antrian didalam buffer. 
``` go
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Rizat"
	channel <- "Software"
	channel <- "Engineer"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Done")
}
```

### Range Channel
Digunakan ketika banyak data yang dikirim, dan kita gak tau berapa banyak channel yang dikirim maka bisa gunakan range untuk mengambil datanya
``` go
func Goroutine_Range(channel chan<- string) {
	for i := 0; i < 10; i++ {
		channel <- "Perulangan ke " + strconv.Itoa(i)
	}
	close(channel)
}

func Test_Goroutine_Range(test *testing.T) {
	channel := make(chan string)

	go Goroutine_Range(channel)

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("Done")
}
```

### Select Channel
Untuk mendapatkan channel yang kita inginkan dari pada menggunakan 1 1 dengan range bisa gunakan `select channel`
``` go
func GiveMeResponse(channel chan<- string) {
	channel <- "Rizat Sakmir"
}

func TestSelectChannel(test *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
```

### Default Select
Ketika select channel dan tidak ada datanya maka akan terus tunggu sampai terjadi error dead lock,  
maka dari itu bisa gunakan default agar tidak terjadi error dead lock,  
dan kita bisa lakukan sesuatu sampai data channel nya ada.
``` go
counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data channel")
		}

		if counter == 2 {
			break
		}
	}
```

### Race Condition
Yaitu dimana dalam 1 waktu bisa terdapat 1000 request yang balapan, yang menjadikan increment/decrement menjadi tidak valid, maka dari itu harus dilakukan locking dari tiap request agar result datanya konsisten. Bisa menggunakan cara tersebut yaitu:
- mutex `race_condition_mutex_test.go` lebih flexible untuk locking
- channel `race_condition_channel_test.go` 
- atomic `race_condition_atomic_test.go` digunakan khusus increment, tidak flexible

### Read Write Locking
Untuk locking Read dan Write dengan Mutex contoh code `rwmutex_test.go`

### Once
Fitur untuk memastikan sebuah function hanya dieksekusi sekali.
``` go
var counter = 0

func CounterIncrement() {
	counter++
}

func OnlyOnce(wg *sync.WaitGroup, once *sync.Once) {
	defer wg.Done()
	once.Do(CounterIncrement)
}

func TestOnce(test *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go OnlyOnce(&wg, &once)
	}

	wg.Wait()
	fmt.Println(counter)
}
```

### Pool
Design pattern bernama object pool pattern, digunakan untuk menyimpan data, untuk menggunakan datanya diambil dari pool, setelah menggunakan bisa dikembalikan datanya kembali ke pool, Implementasi Pool do Go-Lang ini sudah aman dari problem race condition, biasanya digunakan untuk manage data koneksi ke Database.
``` go
func GetDataPool(pool *sync.Pool, wg *sync.WaitGroup) {
	defer wg.Done()
	// Get data pool
	data := pool.Get()

	fmt.Println("This is data pool:", data)
	pool.Put(data)
}

func TestPool(test *testing.T) {
	var pool sync.Pool
	var wg sync.WaitGroup

	// Nilai default pool
	pool.New = func() interface{} {
		return "Default value"
	}

	// Simpan data ke Pool
	pool.Put("Rizat")
	pool.Put("Sakmir")
	pool.Put("Software Engineer")

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go GetDataPool(&pool, &wg)
	}

	wg.Wait()
	fmt.Println("Done")
}
```

### Map
`sync.Map` ini digunakan untuk concurrent menggunakan goroutine
Function yang bisa kita gunakan di Map:
- Store(key, value) - untuk menyimpan data ke Map
- Load(key) - untuk mengambil data dari Map dengan key
- Delete(key) - untuk menghapus data di Map dengan key
- Range(function(keym value)) - untuk melakukan iterasi selutuh data di Map
``` go
func AddToMap(value int, datas *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	datas.Store(value, value)
}

func TestSyncMap(test *testing.T) {
	var datas sync.Map
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go AddToMap(i, &datas, &wg)
	}

	wg.Wait()
	datas.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
```

### Cond
`sync.Cond` atau Condition implementasi locking berbasis kondisi.  
Untuk membuat Cond menggunakan function `sync.NewCond(Locker)`  

### time.Timer
Ketika waktu sudah expire, maka event dikirim ke channel, membuat nya menggunakan `time.NewTimer(duration)`  
``` go
func TestTimer(test *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Time now :", time.Now())

	time := <-timer.C
	fmt.Println("This is time :", time)
}
```

#### time.After()
Jika hanya butuh channel nya saja, tidak membutuhkan data timernya bisa menggunakan function `time.After(duration)`  
``` go
func TestAfter(test *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Time now :", time.Now())

	time := <-channel
	fmt.Println("This is time :", time)
}
```

#### time.AfterFunc()
Untuk menjalankan function dengan delay waktu tertentu, dengan menggunakan `time.AfterFunc()`
``` go
func TestAfterFunc(test *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Execute after 5 second")
		wg.Done()
	})

	wg.Wait()
}
```

### time.Ticker
Kejadian yang berulang, functionya yaitu `time.NewTicker(duration)`, untuk mengehentikan ticker `Ticker.Stop()`
``` go
func TestTimeTicker(test *testing.T) {
	fmt.Println("Run every 1 seconds")
	ticker := time.NewTicker(1 * time.Second)

	looping := 0
	for time := range ticker.C {
		fmt.Println(time)
		if looping == 5 {
			fmt.Println("Ticer di stop")
			ticker.Stop()
			break
		}
		looping += 1
	}
}
```

#### time.Tick()
Mendapatkan channel saja, seperti function Ticker berikut penggunaanya `timer.Tick(duration)`
``` go
func TestTimeTick(test *testing.T) {
	fmt.Println("Run every 1 seconds")
	channel := time.Tick(1 * time.Second)

	looping := 0
	for time := range channel {
		fmt.Println(time)
		if looping == 3 {
			fmt.Println("Tick di stop")
			break
		}
		looping += 1
	}
}
```

### GOMAXPROCS
Untuk mengambil jumlah data thread dan dapat merubah thread, secara default jumlah thread di Go-lang itu sebanyak jumlah CPU yang dikomputer anda, dan kita juga dapat melihat jumlah CPU yang ada dengan `runtime.NumCpu()`  
``` go
func TestGomaxprocs(test *testing.T) {
	total_cpu := runtime.NumCPU()
	fmt.Println("Jumlah CPU :", total_cpu)

	// Rubah jumlah thread
	runtime.GOMAXPROCS(20)
	// Kenapa -1 parameternya karena kalau diatas 0 akan merubah jumlah thread-nya
	total_thread := runtime.GOMAXPROCS(-1)
	fmt.Println("Jumlah Thread :", total_thread)

	total_goroutine := runtime.NumGoroutine()
	fmt.Println("Jumlah Goroutine yang sedang berjalan :", total_goroutine)
}
```