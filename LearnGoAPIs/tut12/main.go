package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            8.99,
		"salad":          5.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping maps

	for key, value := range menu {
		fmt.Printf("key-%v, value-%v\n", key, value)
	}

	// ints as key type

	phonebook := map[int]string{
		239120841: "mario",
		120938109: "luigi",
		189340549: "bowser",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[239120841])

	phonebook[239120841] = "saksham"
	// if the key is not there it will just make a new one and add it
	phonebook[120319209] = "afjds"
	fmt.Println(phonebook)
}
