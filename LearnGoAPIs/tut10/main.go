package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, val := range names {
		initials = append(initials, val[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
func main() {
	fn, sn := getInitials("tifa lockhart")

	fmt.Println(fn, sn)

	fn2, sn2 := getInitials("saksham")
	fmt.Println(fn2, sn2)

}
