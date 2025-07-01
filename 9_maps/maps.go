package main

import "fmt"

func main() {

	// m := make(map[string]string)

	// m["name"] = "Golang"
	// m["area"] = "Backend"

	// fmt.Println(m["name"], m["area"]) // map[name:Golang] //get an element from the map

	m := make(map[string]int)
	m["age"] = 20
	fmt.Println(m["phone"]) // map[phone:0] // if the key does not exist, it returns the zero value of the value type (int in this case, so it returns 0)

	fmt.Println(len(m))
}
