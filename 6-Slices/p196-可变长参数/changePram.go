package main

import "fmt"

func severalStrings(strings ...string) {
	fmt.Println(strings)

}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	severalStrings("1", "s")
	severalStrings("f", "ss", "asd")

	mix(1, true, "asd", "asd")
	mix(12, false, "asd")
}
