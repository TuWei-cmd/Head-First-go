package main

import (
	"fmt"
)

func main() {
	// basic 8: %f %d %s %t %v %#v %T %%
	fmt.Println("GPA", "莫名", "变高！")
	fmt.Printf("A float: %f\n", 89.9)
	fmt.Printf("An interger: %d\n", 1278)
	fmt.Printf("A string: %s\n", "😁")
	fmt.Printf("A boolean: %t\n", true)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Values: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n\n")

	// about %#v
	var x = []string{"1212", "asd"}
	fmt.Printf("x: %#v\n\n", x)

	//the length for %s, %d
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n\n", "Tape", 99)
	//the length for %f
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}
