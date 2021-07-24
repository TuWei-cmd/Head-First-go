package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering up")
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating up")
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = CoffeePot("LuxBrew")
	device.TurnOn()
}
