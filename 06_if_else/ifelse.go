package main

import "fmt"

func main() {

	// age := 18

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// }else{
	// 	fmt.Println("Person is not an adult")
	// }

	// age := 16
	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("person is an teenager")
	// } else {
	// 	fmt.Println("person is an kid")
	// }

	// var role = "admin"
	// hasPermission := true

	// if role == "admin" || hasPermission {
	// 	fmt.Println("yes")
	// }

	//NOTE:  declaring variable inside if

	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager")
	}

	// go does not have ternary
}
