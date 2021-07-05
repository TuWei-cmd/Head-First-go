package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Prince is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars avialble.")
	fmt.Println("Within budgets?", total <= float64(availableFunds))
}
