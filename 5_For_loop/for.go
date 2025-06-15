package main

//for -> only constant in go for looping
import "fmt"

func main() {

	//while LOOP
	// i := 1
	// for i <= 3 {
	// 	fmt.Println("Hello World", i)
	// 	i = i + 1
	// }

	//infinte loop
	// for {
	// 	println("Hello World")
	// }

	//Classic for loop

	// for i := 1; i <= 3; i++ {
	// 	// break // break statement to exit the loop
	// 	if i == 2 {
	// 		continue
	// 	} // continue statement to skip the current iteration
	// 	fmt.Println("Hello World", i)
	// }

	//for range loop
	for i := range 10 {
		fmt.Println("Hello World", i)
	}

}
