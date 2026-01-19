package main

// func printSlice(items []int) {
// 	for _, item := range items {
// 		println(item)
// 	}
// }

// func printSliceString(items []string) {
// 	for _, item := range items {
// 		print(item)
// 	}
// }

func printSliceGeneric[T any](items []T) {
	for _, item := range items {
		println(item)
	}
}

type Stack[T any] struct {
	elements []T
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	printSliceGeneric(numbers)
	strings := []string{"a", "b", "c", "d", "e"}
	printSliceGeneric(strings)

	stack := Stack[int]{elements: []int{10, 20, 30}}
	printSliceGeneric(stack.elements)

	stackStr := Stack[string]{elements: []string{"x", "y", "z"}}
	printSliceGeneric(stackStr.elements)
}
