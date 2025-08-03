package golang_gorutines

import (
	"fmt"
	"sync"
	"testing"
)

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
