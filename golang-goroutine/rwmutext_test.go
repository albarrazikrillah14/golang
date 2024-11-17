package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutext sync.RWMutex
	Balance  int
}

func (a *BankAccount) AddBalance(amount int) {
	a.RWMutext.Lock()
	a.Balance += amount
	a.RWMutext.Unlock()
}

func (a *BankAccount) GetBalance() int {
	a.RWMutext.RLock()
	balance := a.Balance
	a.RWMutext.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.GetBalance())
}
