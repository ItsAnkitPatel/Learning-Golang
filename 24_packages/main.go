package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/itsankitpatel/podcast/auth"
	"github.com/itsankitpatel/podcast/user"
)

func main() {
	auth.LoginWithCredentials("ankit", "1234")

	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "John Doe",
	}
	// fmt.Println("user", user)
	color.Green(user.Email)
}
