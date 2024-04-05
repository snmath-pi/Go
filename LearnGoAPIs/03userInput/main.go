package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome User."
	fmt.Println("Be humble God is the one in command🙏🏻📿")
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type a number: \n")

	// comma ok || error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating us a %v", input)

}
