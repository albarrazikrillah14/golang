package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("Albarra")
	pool.Put("Zikrillah")
	pool.Put("Medomeckz")

	group := sync.WaitGroup{}

	group.Add(1)
	go func() {
		defer group.Done()
		fmt.Println(pool.Get())
		fmt.Println(pool.Get())
		fmt.Println(pool.Get())
	}()

	group.Wait()
	fmt.Println("Selesai")

}
