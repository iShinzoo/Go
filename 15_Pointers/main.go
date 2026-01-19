package main

import "fmt"

// by value
func changeName(num int) {
	num = 20
	fmt.Println("In ChangeNum", num)
}

// by reference
func changeNameRef(nums *int) {
	*nums = 20
	fmt.Println("In ChangeNumRef", *nums)
}

func main() {
	num := 10
	changeName(num)
	fmt.Println("In Main", num)
	fmt.Println("Memory location", &num)

	nums := 30
	changeNameRef(&nums)
	fmt.Println("In Main after ChangeNumRef", nums)
}
