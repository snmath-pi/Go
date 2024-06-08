package main

import "fmt"

func main() {
	age := 40

	// fmt.Println(age <= 50)
	// fmt.Println(age > 40)
	// fmt.Println(age == 40)
	// fmt.Println(age != 100)

	if age < 40 {
		fmt.Println("You're age is less than 40")
	} else if age >= 40 && age <= 50 {
		fmt.Println("You're age is in b/w 40, 50")
	} else {
		fmt.Println("You're age is above 50")
	}

	ages := []int{1, 2, 3, 4, 5, 6}

	for i, val := range ages {
		if i > 2 {
			break
		} else {
			fmt.Println(val)
		}
	}
}
