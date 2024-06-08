package main

import "fmt"

func main() {
	var nameOne = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "bowser"
	nameOne = "Saksham"
	fmt.Println(nameOne, nameThree)

	namenew := "negi"
	fmt.Println(namenew)

	// ints
	var age1 int = 20
	var age2 = 30
	age3 := 100
	fmt.Println(age1, age2, age3)

	// bits
	var numOne int8 = 126
	fmt.Println(numOne)

	var scoreOne float64 = 1232131220.201
	fmt.Println(scoreOne)
}
