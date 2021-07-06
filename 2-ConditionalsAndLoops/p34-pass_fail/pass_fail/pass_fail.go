// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
	// if int(input) == 100 {
	// 	fmt.Println("Perfect")
	// } else if int(input) >= 60 {
	// 	fmt.Println("You pass.")
	// } else {
	// 	fmt.Println("You fail!")
	// }
}
