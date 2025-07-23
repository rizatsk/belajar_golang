package golang_gorutines

import (
	"fmt"
	"sync"
	"testing"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func BalanceSaldo(account *BankAccount, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1; j <= 100; j++ {
		account.AddBalance(1)
		fmt.Println(account.GetBalance())
	}
}

func TestRWMutex(test *testing.T) {
	var wg sync.WaitGroup
	account := &BankAccount{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go BalanceSaldo(account, &wg)
	}

	wg.Wait()
	fmt.Println("Final balance :", account.GetBalance())
}
