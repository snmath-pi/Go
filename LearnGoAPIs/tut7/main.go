package main

import "fmt"

func main() {
	// x := 0

	// for x < 10 {
	// 	fmt.Println(x)
	// 	x += 2
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of x is: ", i, " ")
	// }

	age := []int{1, 2, 3, 4, 5, 6}

	// for i := 0; i < len(age); i++ {
	// 	fmt.Print(age[i], " ")
	// }

	// for index, value := range age {
	// 	fmt.Println(index, value)
	// }

	for _, value := range age {
		fmt.Println(value)
	}
}
