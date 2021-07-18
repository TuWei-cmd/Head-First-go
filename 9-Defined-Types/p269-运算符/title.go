package main

import "fmt"

type Title string

func main() {
	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Alien"))
	fmt.Println(Title("Alien") > Title("Alien"))
	fmt.Println(Title("Alien") + "s")
	// fmt.Println(Title("Jaw 2" - "2"))
}
