// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("average: %.2f\n", sum/sampleCount)
}
