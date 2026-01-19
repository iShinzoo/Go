package auth

import "fmt"

// to export the function as package outside scope use first letter capital
func LoginWithCred(username string, password string) {
	fmt.Println("login user using", username, password)
}
