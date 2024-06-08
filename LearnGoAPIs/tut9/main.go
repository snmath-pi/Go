package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning! %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circArea(r float64) float64 {
	return r * r * math.Pi
}
func pow(x, y int) int {
	res := 1
	for y > 0 {
		if y%2 == 1 {
			res = res * x
		}
		x = x * x
		y >>= 1
	}
	return res
}
func main() {
	sayGreeting("saksham")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	fmt.Println(pow(2, 10))
	fmt.Println(circArea(7))
}
