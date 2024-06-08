package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// to make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
// receiver funcitons
func (b bill) format() string {
	fs := "Bill Breakdown: \n"

	var total float64 = 0

	// list items

	for key, val := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", key+":", val)
		total += val
	}

	fs += fmt.Sprintf("%-25v ...$%v\n", "tip: ", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total: ", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	// (*b) don't need to do thsi for structs in go it's automatic
	b.tip = tip
}
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("temp.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
