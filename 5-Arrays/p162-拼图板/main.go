package main

import "fmt"

func main() {
	numbers := [6]int{12, 312, 123, 123, 1, 12}
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}
