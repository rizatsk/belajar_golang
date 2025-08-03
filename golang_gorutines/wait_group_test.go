package golang_gorutines

import (
	"fmt"
	"sync"
	"testing"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

type TransferService struct {
	SenderUser   *UserBalance
	ReceiverUser *UserBalance
	Amount       int
	Wg           *sync.WaitGroup
}

func Transfer(transfer *TransferService) {
	defer transfer.Wg.Done()

	transfer.SenderUser.Lock()
	transfer.SenderUser.Change(-transfer.Amount)

	transfer.ReceiverUser.Lock()
	transfer.ReceiverUser.Change(transfer.Amount)

	transfer.SenderUser.Unlock()
	transfer.ReceiverUser.Unlock()
}

func TestDeadLock(test *testing.T) {
	var wg sync.WaitGroup
	userRizat := &UserBalance{
		Name:    "Rizat",
		Balance: 100000,
	}

	userHengki := &UserBalance{
		Name:    "Hengki",
		Balance: 100000,
	}

	wg.Add(2)
	go Transfer(&TransferService{
		SenderUser:   userRizat,
		ReceiverUser: userHengki,
		Amount:       10000,
		Wg:           &wg,
	})
	go Transfer(&TransferService{
		SenderUser:   userHengki,
		ReceiverUser: userRizat,
		Amount:       20000,
		Wg:           &wg,
	})

	wg.Wait()
	fmt.Println("User ", userRizat.Name, ", Balance ", userRizat.Balance)
	fmt.Println("User ", userHengki.Name, ", Balance ", userHengki.Balance)
}
