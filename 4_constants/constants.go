package main

import "fmt"

const age = 10 // constants and variables are declared at the package level

func main() {

	const name = "Golang" // string constant
	// const age = 10 // integer constant

	//When using shorthand syntax, it should onl be used in the function scope, shorthand syntax can't be used outside funtion
	// fmt.Println(name)
	// fmt.Println(age)

	const ( // multiple constants
		port = 5000
		host = "localhost"
	)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Port:", port)
	fmt.Println("Host:", host)
}
