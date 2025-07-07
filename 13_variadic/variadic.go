package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total

}

func main() {

	nums := [] int{1,4,8,5,4,3}
	result := sum(nums...)// variadic function can take any number of arguments
	fmt.Println(result)
}
