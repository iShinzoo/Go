package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/iShinzoo/Go/auth"
	"github.com/iShinzoo/Go/user"
)

func main() {
	auth.LoginWithCred("krish", "krish")

	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "krish@gmai.com",
		Name:  "krish",
	}
	fmt.Println(user.Email, user.Name)
	color.Red(user.Email)
	color.Green(user.Name)
}
