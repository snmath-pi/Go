package main

import (
	"fmt"
)

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func add(x1, y1 int) int {
	return x1 + y1
}

func main() {
	fmt.Println(concat("Lauren", "Happy Birthday!"))
	fmt.Println(concat("Elon", "Hope that Tesla thing works out"))
	fmt.Println(concat("Go,", "is fantastic"))

	sendsSofar := 430
	const sendsToAdd = 26
	sendsSofar = incrementSends(sendsSofar, sendsToAdd)
	fmt.Println("you've sent", sendsSofar, "messages")

	firstName, _ := getNames()

	fmt.Println("Welcome to Textiio, ", firstName)

	fmt.Println(YearsUntilEvents(4))
}

func incrementSends(x, y int) int {
	x = x + y
	return x
}

func getNames() (string, string) {
	return "John", "Doe"
}

func YearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrink int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrink = 21 - age
	if yearsUntilDrink < 0 {
		yearsUntilDrink = 0
	}

	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return
}
