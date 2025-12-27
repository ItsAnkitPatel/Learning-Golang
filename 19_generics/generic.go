package main

import "fmt"

// `T any` or `interface{}` means same thing
//
//	func printSlice[T any](items []T) {
//		for _, item := range items {
//			fmt.Println(item)
//		}
//	}

// narrowing down the scope
// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// 'comparable' is another way to write above narrowing down scope generic function for
// including many more types
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// can include multiple generics in a single function
// func printSlice[T comparable, V string](items []T, name V) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// using generics on struct
type stack[T any] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"golang", "typescript"}
	vals := []bool{true, false, true}
	printSlice(nums)
	printSlice(names)
	printSlice(vals)

	myStack := stack[string]{
		elements: []string{"golang"},
	}
	fmt.Println(myStack)
}
