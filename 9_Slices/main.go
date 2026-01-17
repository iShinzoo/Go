package main

import (
	"fmt"
	"slices"
)

func main() {
	// dynamic sized arrays are slices in go

	var nums = make([]int, 0, 3)
	// uintialize slice is nil(null in other languages)

	fmt.Println(nums == nil)

	var names = make([]string, 3, 5)
	names = append(names, "1")
	names = append(names, "")
	fmt.Println(cap(names))
	// capacity (cap) -> maximum number of elements that can fit
	fmt.Println(names)

	// copy function in slices

	nums = append(nums, 9, 9, 3)
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println(nums2)

	// slice operator
	var king = []int{1, 2, 3}
	fmt.Println(king[0:2])
	fmt.Println(king[:1])
	fmt.Println(king[0:])

	// slice package
	var num = []int{5, 3, 8}
	var num1 = []int{7, 3, 4}

	fmt.Println(slices.Equal(num, num1))

	// 2d slices

	matrix := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)

}
