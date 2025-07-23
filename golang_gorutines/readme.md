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
