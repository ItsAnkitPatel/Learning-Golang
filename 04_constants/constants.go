package main

import "fmt"

const age = 30

func main()  {
	// const name string = "golang"
	const name = "golang"
	// name ="javascript" //error
 
	// const (
	// 	port = 5000 
	// 	host = "localhost"
	// )

	// for single line it is separated by ';'
	const (	port = 5000 ;host = "localhost")
	
	println(age)
	fmt.Println(age)
	
}
