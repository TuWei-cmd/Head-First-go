package main

import "fmt"

func main() {
	// change array
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice1)

	// change slice
	array2 := [5]string{"f", "g", "h", "i", "j"}
	slice2 := array2[2:5]
	slice2[1] = "X"
	fmt.Println(array2)
	fmt.Println(slice2)

	// an array, many slices
	array3 := [5]string{"a", "d", "f", "f", "r"}
	slice3 := array3[0:3]
	silce4 := array3[2:5]
	array3[2] = "X"
	fmt.Println(array3)
	fmt.Println(slice3, silce4)
}
