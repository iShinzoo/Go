package main

import (
	"fmt"
	"sync"
)

func task(id int) {
	fmt.Println("doing task", id)
}

func taskk(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing taskk", id)
	// w.Done()
}

// Goroutines are lightweight threads managed by the Go runtime.
// They allow concurrent execution of functions just like kotlin coroutines.
// To start a goroutine, use the 'go' keyword followed by a function call.

func main() {

	// no order in output since goroutines run concurrently
	// for i := 0; i <= 10; i++ {
	// 	go task(i)
	// }

	// time.Sleep(2 * time.Second)

	// using wait groups to wait for goroutines to finish

	var wg sync.WaitGroup

	for j := 0; j <= 5; j++ {
		wg.Add(1)
		go taskk(j, &wg)
	}

	wg.Wait()
}
