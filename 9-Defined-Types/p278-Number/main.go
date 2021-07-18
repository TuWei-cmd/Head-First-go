package main

import "fmt"

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func (n *Number) Double() {
	*n *= 2
}

func main() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(42)
	four := Number(13)
	four.Add(89)
	four.Subtract(1290)
}
