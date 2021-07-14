package main

import "fmt"

func main() {
	// empty elements in the slice
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Printf("%#v\n%#v\n", floatSlice, boolSlice)

	// nil silces
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v \nstringSlice: %#v \n", intSlice, stringSlice)

	// append nil slice
	intSlice = append(intSlice, 12)
	fmt.Printf("intSlice: %#v\n", intSlice)

	// great
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("slice: %#v\n", slice)
}
