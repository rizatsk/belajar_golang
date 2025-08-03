package golang_gorutines

import (
	"fmt"
	"runtime"
	"testing"
)

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
