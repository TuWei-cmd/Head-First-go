package main

import "fmt"

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 0.001)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}
func (m Milliliters) ToLiters() Liters {
	return Liters(m * 1000)
}

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}
func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}
func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%s = %s\n", soda, soda.ToGallons())
	water := Milliliters(1323)
	fmt.Printf("%s = %s\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%s = %s\n", milk, milk.ToLiters())
	fmt.Printf("%s = %s\n", milk, milk.ToMilliliters())
}
