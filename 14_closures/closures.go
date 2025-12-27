package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count += 1
		return count
	}
}
func main() {
	//NOTE: interesting thing to notice here, in terms of memory
	// increment := counter

	// fmt.Println(increment())
	// fmt.Println(increment())
	
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())

}
