package main

import (
	"fmt"
	"slices"
)

// slice means dynamic array
func main() {
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(nums == nil) //true
	// fmt.Println(len(nums))

	var nums = make([]int, 0, 5)
	fmt.Println(nums)
	fmt.Println(nums == nil) //false
	// capacity => maximum numbers of elements can fit but it will be changed in slice since it's dynamic
	fmt.Println(cap(nums))

	// append add elements at last
	nums = append(nums, 1, 2, 3, 4)

	fmt.Println(nums)
	fmt.Println("current length", len(nums))
	fmt.Println("new capacity", cap(nums))
	fmt.Println(nums[2])

	// copy function

	var nums2 = make([]int, len(nums))

	copy(nums2, nums)

	fmt.Println("copied array", nums2)

	// slice operator

	var nums3 = []int{1, 2, 3, 4, 5}

	fmt.Println("slicing", nums3[0:2])
	fmt.Println("slicing", nums3[:2]) //same as above
	fmt.Println("slicing", nums3[1:])

	// slices package
	var nums4 = []int{1, 2}
	var nums5 = []int{1, 2}

	fmt.Println("both array is equal", slices.Equal(nums4, nums5))

	// 2D slices

	var nums2DSlice = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D Slice", nums2DSlice)
}
