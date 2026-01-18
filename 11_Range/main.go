package main

import "fmt"

// Range -> used to iterate over data structures
func main() {
	nums := []int{10, 20, 30, 40, 50}
	sum := 0

	for i, v := range nums {
		fmt.Println(i, v)
		sum += v
	}

	fmt.Println("Sum:", sum)

	// iterating maps using range
	maps := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}
	for k, v := range maps {
		fmt.Println(k, v)
	}

	// iterating over string using range gives unicode code point rune values or ascii values
	// i -> index, c -> rune (character)
	for i, c := range "hello" {
		fmt.Println(i, c, string(c))
	}
}
