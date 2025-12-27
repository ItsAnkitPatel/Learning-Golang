package main

import "fmt"

// passing n number of parameters is called variadic functions in golang
// for passing any type we use `interface{}` example: func sum(nums ...interface{})
func sum(nums ...int) int {
	total := 0

	for _, nums := range nums {
		total += nums
	}
	return total
}

func main() {

	nums := []int{3, 4, 5, 6}
	// result := sum(3, 4, 5, 6)
	
	// spreading slices using three dots
	fmt.Println("result", sum(nums...))
}
