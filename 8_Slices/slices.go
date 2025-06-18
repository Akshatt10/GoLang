package main

import "fmt"

//slices are dynamic arrays and most used data structure in go

func main() {

	//uninitalized slice is "NILL"
	// var nums []int

	// fmt.Println(nums)
	// fmt.Println(len(nums)) // length of slice

	var nums = make([]int, 2)
	//capacity the is the maximum number of elements the slice can hold
	fmt.Println(nums, cap(nums)) // capacity of slice
	fmt.Println(nums, len(nums)) // length of slice
}
