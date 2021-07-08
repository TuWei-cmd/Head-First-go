package main

import "fmt"

func double(myPointer *int) {
	*myPointer *= 2
}

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}
