package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestHelloWorld(t *testing.T) {
	RunHelloWorld()

	fmt.Println("Done")

	/*
		Hello World
		Done
	*/
}

func TestHelloWorldGoroutine(t *testing.T) {
	go RunHelloWorld() // berjalan secar asynchronus

	fmt.Println("Done")

	time.Sleep(1 * time.Second)

	// kadang fungsi sudah selesai tapi fungsi goroutine belum dipanggil, jadi fungsi selesai dan fungsi goroutine tidak dipanggil
	/*
		Done
		HelloWorld
	*/
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	//terurut
	// for i := 0; i < 10000; i++ {
	// 	DisplayNumber(i)
	// }

	// tidak terurut
	for i := 0; i <= 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
