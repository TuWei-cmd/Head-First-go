// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if grade == 100.0 {
		fmt.Println("Perfect")
	} else if grade >= 60.0 {
		fmt.Println("You pass.")
	} else {
		fmt.Println("You fail!")
	}
}
