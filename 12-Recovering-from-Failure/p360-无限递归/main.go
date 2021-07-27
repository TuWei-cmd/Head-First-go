package main

import "fmt"

func recurses() {
	fmt.Println("Oh on, I'm stuck")
	recurses()
}

func main() {
	recurses()
}
