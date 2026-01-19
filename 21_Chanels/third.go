package main

import (
	"fmt"
	"time"
)

// buffered channel

func emialSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		// send email
		println("Sending email to ", email)
		time.Sleep(time.Second)
	}
}

// receive only channel
// func emialSender(emailChan <-chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		// send email
// 		println("Sending email to ", email)
// 		time.Sleep(time.Second)
// 	}
// }

// send only channel
// func emialSender(emailChan <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		// send email
// 		println("Sending email to ", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emialSender(emailChan, done)

	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending...")

	// // this is important to close the channel otherwise deadlock occurs
	// close(emailChan)
	// <-done

	// multiple channels

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received from chan1: ", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received from chan2: ", chan2Val)
		}
	}

}
