package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)

	/*
		counternya != 1000 * 100, karena ada beberapa looping terjadi race condition
	*/
}

func TestMutext(t *testing.T) {
	counter := 0
	var mutext sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutext.Lock()
				counter = counter + 1
				mutext.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", counter)

}
