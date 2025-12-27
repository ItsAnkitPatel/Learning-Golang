package main

import "fmt"

func changeNum(num *int) {
	// deference
	*num = 5
	fmt.Println("in changeNum", *num)
}
func main() {
	num := 1
	fmt.Println("memory address", &num)
	changeNum(&num)
	fmt.Println("after changeNum in main", num)
}
