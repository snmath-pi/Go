package main

import "fmt"

func main() {
	fmt.Print("hello, ")
	fmt.Print("world! \n")

	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas")
	age := 20
	name := "Saksham"

	fmt.Println("my name is, ", name, " and my age is, ", age)

	// Printf() formatted string

	fmt.Printf("my age is %v and my name is %v", age, name)

	// %q -> strings

	// %f -> decimal
	// %T -> type of varible
}
