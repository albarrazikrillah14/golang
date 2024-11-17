package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronus(g *sync.WaitGroup) {
	defer g.Done()

	g.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		// group := sync.WaitGroup{}

		// for i := 0; i <= 100; i++ {
		// 	go RunAsynchronus(&group)
		// }

		// group.Wait()
	})

	t.Run("Test 2", func(t *testing.T) {
		group := sync.WaitGroup{}

		counter := 0
		var mutext sync.Mutex

		for i := 0; i < 1000; i++ {
			for j := 0; j < 100; j++ {
				group.Add(1)
				go func() {
					defer group.Done()
					mutext.Lock()
					counter++
					mutext.Unlock()
				}()
			}
		}

		group.Wait()
		fmt.Println(counter)
	})
}
