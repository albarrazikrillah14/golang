package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go SendString("Test", channel)
	go SendString("Halo", channel)

	// get from channel
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel) // menunggu data, namum tidak ada
}

func SendString(value string, c chan string) {
	c <- value
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Albarra Zikrillah"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(time.Second * 2)
	channel <- "Albarra Zikrillah"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestOnyInOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(time.Second * 2)
}

func TestBufferedChannel(t *testing.T) {
	/*
		tidak akan melakukan bloking jika channel dalam bentuk buffered, namun jika data yang dimasukkan melebihi ukuran channel maka akan melakukan bloking, sampai data yang didalam channel ada yang mengambil
	*/
	channel := make(chan string, 3)

	channel <- "Albarra"
	channel <- "Zikrillah"
	channel <- "Medomeckz"
	// channel <- "Uhuyy" // terjadi bloiking, dead lock

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}

func TestChannelRange(t *testing.T) {
	channel := make(chan string)

	go func() {
		index := 0
		for {

			if index == 10 {
				close(channel)
			}

			channel <- "Perulangan ke " + strconv.Itoa(index)
			index += 1
		}
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")

		}

		if counter == 2 {
			break
		}
	}
}
