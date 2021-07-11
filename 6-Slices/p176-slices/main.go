package main

import "fmt"

func main() {
	// 声明slices
	var myArray [3]string
	var mySlice []string
	fmt.Printf("%#v  %#v\n", myArray, mySlice)
	fmt.Println(len(mySlice))

	var notes []string = make([]string, 7)
	notes[1] = "re"
	notes[0] = "do"
	notes[2] = "mi"

	primes := make([]int, 5)
	primes[0] = 21
	primes[3] = 12

	// len of slices
	fmt.Printf("%s	v:%#v T:%T len:%v\n", "mySlice", mySlice, mySlice, len(mySlice))
	fmt.Printf("%s	v:%#v T:%T len:%v\n", "notes", notes, notes, len(notes))
	fmt.Printf("%s	v:%#v T:%T len:%v\n", "primes", primes, primes, len(primes))

	// 字面量
	notes2 := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Printf("%#v\n", notes2)
}
