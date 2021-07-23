package main

import (
	"fmt"
	"geo"
)

func main() {
	location := geo.Coordinates{
		Latitude:  37.42,
		Longitude: -122.08,
	}
	location.SetLatitude(38.2)
	location.SetLongitude(-123.4)
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
