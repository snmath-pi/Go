package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err

}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new Bill Name: ", reader)
	b := newBill(name)
	fmt.Println("Bill Created-", b.name)
	return b

}

func promptOptions(b bill) {
	r := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - additem, s - save bill, t - add tip): ", r)
	switch opt {
	case "a":
		fmt.Println("you chose a")
		name, _ := getInput("Enter the item name", r)
		price, _ := getInput("Item Price: ", r)
		p, _ := strconv.ParseFloat(price, 64)
		b.addItem(name, p)
		fmt.Println("item added", name, price)
		promptOptions(b)
	case "t":
		fmt.Println("you chose t")
		tip, _ := getInput("Enter tip amount ($): ", r)
		u, _ := strconv.ParseFloat(tip, 64)
		b.updateTip(u)
	case "s":
		fmt.Println("you chose s")
		b.save()
	default:
		fmt.Println("you chose invalid option", opt)
		promptOptions(b)
	}
}
func main() {

	mybill := createBill()

	promptOptions(mybill)
}
