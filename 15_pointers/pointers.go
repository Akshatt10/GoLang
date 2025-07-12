package main

import "fmt"

//by reference function

func byReference(num *int) {
	*num = 10
	fmt.Println("Inside byReference function:", *num)
}

func main() {
	num := 1

	byReference(&num)
	fmt.Println(num)
}
