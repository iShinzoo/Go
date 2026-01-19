package main

import "fmt"

// using channels as wait groups

// in case of multiple goroutines use wait groups
// in case of single goroutine use channels

// goroutine synchronizer
func task(done chan bool) {

	defer func() {
		done <- true
	}()

	fmt.Println("Task is being done")
}

func main2() {

	done := make(chan bool)
	go task(done)
	<-done //block

}
