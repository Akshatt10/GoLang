package main

import (
	"fmt"
	"slices"
)

//slices are dynamic arrays and most used data structure in go

func main() {

	//uninitalized slice is "NILL"
	// var nums []int

	// fmt.Println(nums)
	// fmt.Println(len(nums)) // length of slice

	// var nums = make([]int, 2,5)
	//capacity the is the maximum number of elements the slice can hold
	// fmt.Println(nums, cap(nums)) // capacity of slice
	// fmt.Println(nums, len(nums)) // length of slice

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums) //copying slice

	// fmt.Println(nums, nums2) // capacity of slice

	//SLICE OPERATOR

	// var nums = []int{1, 2, 3}

	// fmt.Println(nums[0:2]) // slicing the slice, it will return a new slice with elements from index 0 to 2 (exclusive)
	// fmt.Println(nums[:2])  // slicing the slice, it will return a new slice with

	//Slice Package

	var nums = []int{1, 2, 3}
	var num2 = []int{4, 5, 6}

	fmt.Println(slices.Equal(nums, num2)) // checking if two slices are equal

}
