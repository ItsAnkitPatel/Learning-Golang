package auth

import "fmt"

// if we want the function to be accessible outside of the package we start the function name with Capital letter
func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}
