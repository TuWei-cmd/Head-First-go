package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calender"
)

func main() {
	date := calender.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
