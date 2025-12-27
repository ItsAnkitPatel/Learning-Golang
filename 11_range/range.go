package main

import "fmt"

// range -> for iterating over data structures
func main() {
	nums := []int{6, 7, 8, 9, 10}

	// normal iteration of slices using classic for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	sum := 0
	// using range for iterating slices
	for i, num := range nums {
		sum += num
		fmt.Println("index:", i, "value:", num)
	}
	fmt.Println("sum", sum)

	// using range for iterating map
	m := map[string]string{"fname": "john", "lname": "doe"}

	// for getting key and value both
	// for k, v := range m {
	// 	fmt.Println("key:", k, "value:", v)
	// }

	// if want only keys then use this syntax
	for k := range m {
		fmt.Println(k)
	}

	// using range over string

	for i, c := range "golang" {
		fmt.Println("index:", i, "char:", c) //unicode will be printed
		fmt.Println("index:", i, "char:", string(c))
	}
}
