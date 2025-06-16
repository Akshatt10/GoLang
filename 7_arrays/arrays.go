package main

import "fmt"

func main() {
	var nums [4]int

	//adding values to the array
	nums[0] = 1
	nums[1] = 2

	//getting element at index 0
	fmt.Println("Element at index 0:", nums[0])
	//length of array
	fmt.Println("Length of nums:", len(nums))

	//printing the entire array
	fmt.Println("Array nums:", nums)
}
