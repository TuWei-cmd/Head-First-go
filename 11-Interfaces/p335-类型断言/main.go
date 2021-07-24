package main

import "fmt"

type Robot string

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}
func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

type NoiseMaker interface {
	MakeSound()
}

func main() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}
