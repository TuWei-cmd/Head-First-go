package main

import "fmt"

type Liters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 0.001)
}

type Milliliters float64

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}
func (m Milliliters) ToLiters() Liters {
	return Liters(m * 1000)
}

type Gallons float64

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}
func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(1323)
	fmt.Printf("%.3f Milliliters equals %0.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons %0.3f Milliliters\n", milk, milk.ToMilliliters())
}
