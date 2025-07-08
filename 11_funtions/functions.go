package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func getlanguages() (string, string, string){
// 	return "Go", "Python", "JavaScript"
// }


//funtion inside a function

func process(fn func (a int) int) {
	fn(1)
}

func main() {

	// result := add(3, 5)
	// fmt.Println(result)

	// result1, result2, _ :=  getlanguages()
	// fmt.Println(result1, result2)

	fn := func(a int) int {
		fmt.Println("Processing:", a)
		return a * 2
	}
	process(fn)
}
