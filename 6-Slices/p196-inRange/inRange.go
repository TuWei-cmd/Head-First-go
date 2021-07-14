package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, -12, 12, 13, 13, 124, 12))
	fmt.Println(inRange(129, 23123, 12, 8123480129, 12831299, 12839, 18300))
}
