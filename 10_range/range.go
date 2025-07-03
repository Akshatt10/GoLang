package main

import "fmt"

//iterating over data structures in Go
func main() {
	// nums := []int{1, 2, 3, 4, 5}

	// // for i:= 0; i<=3; i++{
	// // 	fmt.Println(nums[i])
	// // }
	// // sum := 0
	// for i, num := range nums {

	// 	fmt.Println(num, i)

	// }
	// fmt.Println(sum)

	//iterating the map

	// m := map[string]string{"fname": "John", "age": "30", "city": "New York"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	//iterating over a string
	for i, c := range "golang" {
		// fmt.Println(i, c) // i is the index, c is the rune (character)
		fmt.Println(i, string(c))
	}

}
