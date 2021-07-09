package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 8
	fmt.Printf("%.2f degree Celius\n", celsius)
}
