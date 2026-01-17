package main

import "fmt"

func main() {

	var nums [5]int

	nums[0] = 10
	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)

	var names [3]string
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	fmt.Println(len(names))
	fmt.Println(names)

	// zero value of array elements
	// int -> 0, string -> "", bool -> false

	numss := [3]int{1, 2, 3}

	fmt.Println(numss)

	// 2d array
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrix)

	// in fixed size arrays are used, that is predictable
}
