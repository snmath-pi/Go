package main

import "fmt"

func main() {
	// fmt.Println("Hello World")

	// go runs faster than most interprated lang and faster compilation than c++, java etc
	// go executes slower than rust
	fmt.Println("textio server is starting.")

	var username string = "Saksham"
	// var password int = 12345
	var password string = "12345"
	// fmt.Println("Authorization: Basic", username+":"+password) // throws error mismatched types
	// compile time error
	fmt.Println("Authorization: Basic", username+":"+password)

}
