package main

import "fmt"

func change(x string) string {
	x = "wedge"
	return x
}
func main() {
	name := "tifa"
	name = change(name)
	fmt.Println(name)
}
