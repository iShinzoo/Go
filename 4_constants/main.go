package main

import "fmt"

const age = 25

func main() {

	const name = "krish"

	// shorthand declaratiion canot be used outside functions

	// multiple constants
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(port, host)
}
