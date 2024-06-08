package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "hello there friends"
	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))

	// fmt.Println(strings.Index(greeting, "llo"))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println(greeting)

	ages := []int{324, 131, 1321, 1, 10, 2}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 131)
	fmt.Println(index)

	names := []string{"saksham", "negi", "maggi", "momos"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "saksham"))
}
