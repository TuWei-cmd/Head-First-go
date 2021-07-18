package main

import "fmt"

type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a")
	pointer := &value

	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()

	// warning!!!
	pointer2 := &MyType("a")
	MyType("asd").pointerMethod()
	// Save it in a value
	value2 := MyType("asd")
	value2.pointerMethod()
}
