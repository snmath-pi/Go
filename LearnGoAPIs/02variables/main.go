package main

import "fmt"

// if first letter is capital letter then that variable is public
const LoginToken string = "ghabbhid"

func main() {
	// fmt.Println("Variables")
	var username string = "Saksham"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T. \n", isLoggedIn)

	var val uint8 = 255
	fmt.Println(val)

	var floatVal float32 = 54.34328943289489327
	fmt.Println(floatVal)
	fmt.Printf("Variable is of the type: %T \n", floatVal)

	// default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of the type: %T \n", anotherVariable)

	// implicit type of declaring the variables

	var website = "sakshamnegi.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 10
	fmt.Println(numberOfUser)

	// fmt.Println(LoginToken)
	// fmt.Printf("Variable is of the type: %T \n", LoginToken)

}
