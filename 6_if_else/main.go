package main

import "fmt"

func main() {

	// we can declare and initialize a variable in the if statement

	if age := 20; age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age < 65 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior Citizen")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" && hasPermission {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}

	// go does not have ternary operator
}
