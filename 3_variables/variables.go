package main

import "fmt"

func main() {
	// Variables
	// Declare a variable

	//var name string = "golang"  // string variable

	//go automatically infers the type
	var name = "JOHN DOE" // string variable

	var age int = 30 // integer variable

	fmt.Println(name)

	fmt.Println(age)

	//Shorthand syntax

	test := "Jane Doe" // string variable
	testt := 25        // integer variable

	fmt.Println(test)
	fmt.Println(testt)

	// var price float32 = 50.5 // float variable
	var price = 50.5 // float variable
	// go automatically infers the type
	// price := 50.5 // float variable
	// shorthand syntax
	fmt.Println(price)

}
