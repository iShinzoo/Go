package main

import "fmt"

// variadic function can take n no. of inputs

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// to recive any type of inputs we can use empty interface instead of ...int to ...interface{}

func main() {
	fmt.Println(1, 2, 3, 4, 5, "hello") // can have n no. of inputs -> variadic function

	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println("Sum is:", result)
}
