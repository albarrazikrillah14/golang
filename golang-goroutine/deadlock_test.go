package golang_goroutines

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

func (u *UserBalance) GetBalance() int {
	u.Mutex.Lock()
	balance := u.Balance
	u.Mutex.Unlock()
	return balance
}

func (u *UserBalance) AddBalance(amount int) {
	u.Mutex.Lock()
	u.Balance += amount
	u.Mutex.Unlock()
}

func (u *UserBalance) TransferTo(u2 *UserBalance, amount int) {
	u.AddBalance(-amount)
	u2.AddBalance(amount)
}

func TestDeadLock(t *testing.T) {

	user1 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Albarra",
		Balance: 100000,
	}

	user2 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Medomeckz",
		Balance: 200000,
	}

	user2.TransferTo(&user1, 100000)

	fmt.Println(user1.Name, " ", user1.Balance)
	fmt.Println(user2.Name, " ", user2.Balance)

}

/*

Dead lock terjadi ketika 2 Mutex saling Lock

misalkan ada 2 fungsi(A, B) dan 2 User (U1, U2)

	A.U1 Lock
	A.U2 Lock

	B.U2 Lock
	B.U1 Lock

	A.U2 UnLock
	A.U1 UnLock

	B.U2 UnLock
	B.U1 UnLock
*/
