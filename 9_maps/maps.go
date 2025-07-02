package main

import (
	"fmt"
	"maps"
)

func main() {

	// m := make(map[string]string)

	// m["name"] = "Golang"
	// m["area"] = "Backend"

	// fmt.Println(m["name"], m["area"]) // map[name:Golang] //get an element from the map

	// m := make(map[string]int)
	// m["age"] = 20
	// fmt.Println(m["phone"]) // map[phone:0] // if the key does not exist, it returns the zero value of the value type (int in this case, so it returns 0)

	// fmt.Println(len(m))

	// delete(m, "age") // delete an element from the map
	// fmt.Println(m)   // map[] // after deletion, the map is empty

	// clear(m) // clear the map, it will be empty
	// fmt.Println(m)

	// fmt.Println(m) // map[age:20 salary:100000 phone:1234567890]

	// m := map[string]int{"age": 20, "salary": 100000, "phone": 1234567890}

	// //checking whether element prsent in the map or not

	// v, ok := m["age"]

	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("Key 'age' exists in the map")

	// }else {
	// 	fmt.Println("Key 'age' does not exist in the map")

	m1 := map[string]int{"age": 20, "salary": 100000, "phone": 1234567890}
	m2 := map[string]int{"age": 20, "salary": 100000, "phone": 1234567890}

	// fmt.Println(m1 == m2) // false // maps are not comparable, so we cannot compare two maps directly
	// to compare two maps, we need to use the maps package from the standard library
	fmt.Println(maps.Equal(m1, m2)) // true // this will compare
}
