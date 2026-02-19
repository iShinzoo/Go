package main

import "fmt"

func main() {

	// initial channel with data
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))
	for _, d := range data {
		input <- d
	}
	close(input)

	// first stage of pipeline -> 2x input value
	doubleOutput := make(chan int)
	go func() {
		defer close(doubleOutput)
		for num := range input {
			doubleOutput <- num * 2
		}
	}()

	// second stage of pipeline -> ^2 doubled values
	squareOutput := make(chan int)
	go func() {
		defer close(squareOutput)
		for num := range input {
			squareOutput <- num * num
		}
	}()

	// third stage -> print squared values
	for result := range squareOutput {
		fmt.Println(result)
	}
}
