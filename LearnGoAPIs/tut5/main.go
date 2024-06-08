package main

import "fmt"

func main() {
	// var ages [5]int = [5]int{1, 2, 3, 4, 5}
	var ages = [3]int{1, 2, 3}
	names := [4]string{"yoshi", "mario", "luigi", "bowser"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices arrays under the hood
	score := []int{1, 2, 3, 4}
	fmt.Println(score[0], len(score))
	score = append(score, 5)
	fmt.Println(score)

	fmt.Println(score[2:3])
	fmt.Println(score[:])

}
