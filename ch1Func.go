package main

import "fmt"

func main() {
	var Price = 100
	fmt.Println("price of the shoes is", Price ,"dollars")

	var taxRate = 0.84
	var tax = float64(Price) * taxRate
	fmt.Println("tax to be paid is", tax ,"dollars")

	var total float64 = float64(Price) + tax
	fmt.Println("Total cost to be paid is", total ,"dollars")

	var availableFunds = 120
	fmt.Println(availableFunds ,"dollars available")
	fmt.Println("within budget?", total <= float64(availableFunds))
}

