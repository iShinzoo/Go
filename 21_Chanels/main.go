package main

import (
	"fmt"
	"time"
)

// Channels are a way for goroutines to communicate with each other and synchronize their execution.
// They can be thought of as pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values in another goroutine.

// sending data
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

// for receiving data
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

// <- operator is used to send and receive data from channels.

func main1() {

	// messageChan := make(chan string)

	// messageChan <- "Hello, Channels!" // This will cause a deadlock because there is no goroutine receiving from the channel.

	// msg := <-messageChan

	// fmt.Println(msg)

	// for sending
	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// for receiving
	result := make(chan int)
	go sum(result, 5, 10)
	res := <-result
	fmt.Println("Sum is", res)

}
