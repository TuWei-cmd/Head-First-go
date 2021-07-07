package main

import (
	"fmt"
	"log"
	"math"
)

func squereRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get squre root of negative number %#v", number)
	}
	return math.Sqrt(number), nil
}

func main() {
	root, err := squereRoot(-12.0908)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.3f", root)
}
