package main

import (
	"fmt"
	"time"
)

// if we want to communicate between goroutines we can use channels
// channels are blocking operation

// retrieving data from channel
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

// sending data to channel
func sum(result chan int, num1, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}
func main() {
	// creating channel
	// messageChan := make(chan string) //NOTE: these are called `Unbuffered channel` means this version is blocking type

	// // sending data inside channel
	// messageChan <- "ping" //blocking

	// // receiving data from channel
	// msg := <-messageChan
	// fmt.Println(msg)

	// numChan := make(chan int)
	// go processNum(numChan)

	// // numChan <- 5
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 1)

	// result := make(chan int)
	// go sum(result, 4, 5)

	// res := <-result //blocking, that's why we don't need to use time sleep
	// fmt.Println("res", res)

	// we can do similar work we did via sync.WaitGroup
	// done := make(chan bool)
	// go task(done)

	// <-done //block, until data is received, so no need to use time.sleep

	//NOTE: buffered channels, they are non-blocking but there is a catch, for the below example like for the size of 100 the channel will be non-blocking but after 100 it will be blocking, this is interesting
	// emailChan := make(chan string, 100)

	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := range 5 {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending to channel.")
	//IMPORTANT: it is important to close channels to avoid error
	// close(emailChan)

	// <-done

	// Listening to multiple channels

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}
}
