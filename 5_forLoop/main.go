package main

import "fmt"

func main() {

	// while loop style
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	println(1)
	// }

	// classic for loop
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// range based for loop
	for k := range 3 {
		fmt.Println(k)
	}

}
