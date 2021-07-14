package main

import "fmt"

func average(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	fmt.Println(average(12, 32, 3, 213, 21, 3, 21, 3, 3))
	fmt.Println(average(123, 124, 125, 346, 346374, 6346, 346, 436, 346, 34))
}
