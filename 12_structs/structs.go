package main

import "fmt"

func main() {

	Name := User{"Akshat", 20, true, "akshat@gmail.com"}
	fmt.Println(Name)
	fmt.Printf("Name of the user is %+v\n", Name.Name)

}

type User struct {
	Name     string
	Age      int
	verified bool
	email    string
}
