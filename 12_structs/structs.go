package main

import (
	"fmt"
	"time"
)

// Address struct can be used inside other structs
type Address struct {
	Street string
	City   string
}

// User struct with an embedded Address struct
type User struct {
	Name      string
	Age       int
	verified  bool
	email     string
	createdAt time.Time
	Address   Address // Embedding the Address struct
}

// Method with a value receiver (User) - doesn't modify the original struct
func (u User) printDetails() {
	fmt.Printf("Name: %s, Age: %d, Email: %s, Verified: %t\n", u.Name, u.Age, u.email, u.verified)
}

// Method with a pointer receiver (*User) - MODIFIES the original struct
func (u *User) verify() {
	u.verified = true
}

// "Constructor" function to create and initialize a new user
func NewUser(name, email string, age int) *User {
	return &User{
		Name:      name,
		Age:       age,
		email:     email,
		verified:  false, // Start as unverified by default
		createdAt: time.Now(),
	}
}

func main() {
	// 1. Create a user using the constructor function
	user := NewUser("Akshat", "akshat@gmail.com", 20)

	// 2. Add address details to the embedded struct
	user.Address = Address{
		Street: "123 Go Lane",
		City:   "Dev City",
	}

	fmt.Println("--- Initial User State ---")
	user.printDetails() // Call the method
	fmt.Printf("Address: %s, %s\n", user.Address.Street, user.Address.City)

	// 3. Modify the user with a pointer-receiver method
	fmt.Println("\n--- Verifying User ---")
	user.verify() // This will modify the original 'user' struct

	fmt.Println("\n--- Final User State ---")
	user.printDetails()
}
