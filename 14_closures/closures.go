package main

import "fmt"

func counter() func() int {

	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {

	increment := counter()   // create a closure
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

}
