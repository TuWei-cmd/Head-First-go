package main

import (
	"html/template"
	"log"
	"os"
)

type Invoice struct {
	Name     string
	Paid     bool
	Chargers []float64
	Total    float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	html, err := template.ParseFiles("bill.html")
	check(err)
	bill := Invoice{
		Name:     "Mary Gibbs",
		Paid:     true,
		Chargers: []float64{23.19, 1.13, 24.5},
		Total:    67.88,
	}
	err = html.Execute(os.Stdout, bill)
	check(err)
}
