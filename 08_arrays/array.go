package main

import "fmt"

// NOTE: array is numbered sequence of specific length
// if you already knew the array then use fixed size array for memory optimization
// if you don't know coming array size then use Slices
func main() {
	var nums [4]int

	fmt.Println("array length", len(nums))

	nums[0] = 1

	fmt.Println("value at index 0 is", nums[0])

	fmt.Println("all value", nums)

	var boolArr [4]bool

	fmt.Println("bool array", boolArr)

	var name [3]string
	name[0] = "golang"
	fmt.Println("name array", name)

	// to define array in single line
	numArr := [3]int{1, 2, 3}
	fmt.Println("num array", numArr)

	// 2D arrays

	num2DArr := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println("2D array", num2DArr)

}
