package main

import "fmt"

// name:="Ankit" //error
var name ="Ankit"
var anotherName string = "Yash"

func main(){
	// var name string = "golang"

	// short type, infer type automatically
	var name = "golang"

	// var isAdult bool = true
	var isAdult = true

	// var age int = 23
	var age = 23

	// shorthand syntax
	firstName := "Ankit"


	var lastName string
	lastName = "Patel"

	// var price float32 = 50.5
	// var price = 50.5
	 price := 50.5


	fmt.Println(name)
	fmt.Println(isAdult)
	fmt.Println(age)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(price)
}
