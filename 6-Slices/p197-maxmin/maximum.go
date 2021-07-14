package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(maximum(123, 123, 124, 125, 45, 345))
	fmt.Println(maximum(123, 123, 1204, 125, 45, 345))
}
