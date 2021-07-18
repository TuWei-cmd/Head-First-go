package main

import (
	"github.com/headfirstgo/magazine"
)

func main() {
	var s = magazine.Subscriber{
		Name:   "Corvette",
		Rate:   4.99,
		Active: true,
	}
	magazine.PrintInfo(&s)
}
