package main

import (
	"fmt"
)

func main() {
	// i := 5
	// NOTE: in golang we don't need to write 'break' after every case

	// simple switch
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// Multiple condition switch

	// switch time.Now().Weekday() {
	// // if it's Saturday or Sunday
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }

	// Type switch
	whoAmI := func(i interface{}) {
		// switch i.(type){ //this will also work
		switch t := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's an string")
		case bool:
			fmt.Println("it's a bool")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI("golang")
	whoAmI(true)
	whoAmI(5)
	whoAmI(5.6)

	// For checking type using fmt
	fmt.Printf("%T", 4)
}
