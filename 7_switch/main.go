package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three or Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	// multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm an integer")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println("I'm of a different type", t)
		}
	}

	whoAmI(42)
}
