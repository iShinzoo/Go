package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int, 10)

	// creating 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	// send 10 jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	time.Sleep(3 * time.Second)
}
