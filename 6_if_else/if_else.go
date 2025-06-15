package main

import "fmt"

func main() {

	// age := 34

	// if age < 18 {
	// 	fmt.Println("You are a minor")
	// } else if age >= 18 && age < 60 {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are a senior citizen")
	// }

	// var role = "admin"
	// var hasPermissions = true

	// if role == "admin" && hasPermissions {
	// 	fmt.Println("You have admin access")
	// }

	if age := 34; age < 18 {
		fmt.Println("You are a minor")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a senior citizen")
	}

	//go does not support ternary operator like other languages

}
