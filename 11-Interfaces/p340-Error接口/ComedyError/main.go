package main

import "fmt"

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	var err error
	err = ComedyError("What's a programmmer's facorite beer? Logger!")
	fmt.Println(err)
}
