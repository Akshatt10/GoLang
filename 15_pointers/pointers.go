package main

import "fmt"

func changeValue(num *int) {
	*num = *num + 1
}

func main(){
	num:= 1
	changeValue(&num)
	fmt.Println(num) // Output: 2
}