// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(input)
}
