package golang_gorutines

import (
	"fmt"
	"sync"
	"testing"
)

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
