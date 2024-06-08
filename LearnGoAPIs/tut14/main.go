package main

import "fmt"

func change(x *string) {
	*x = "wedge"
}
func main() {
	name := "saksham"
	fmt.Println("pointer to the name is", &name)

	m := &name

	fmt.Println(m)
	fmt.Println("value: ", *m)

	change(m)

	fmt.Println(name)

	name2 := "mario"
	fmt.Println(name2)
	change(&name2)
	fmt.Println(name2)

}
