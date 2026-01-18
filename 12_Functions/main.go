package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, int, bool) {
	return "goLang", "python", 2, true
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	result := add(2, 6)
	fmt.Println(result)
	fmt.Println(getLanguages())

	fn := processIt()
	fmt.Println(fn(6))
}
