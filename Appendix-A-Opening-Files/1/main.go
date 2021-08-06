package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	options := os.O_RDWR | os.O_CREATE
	file, err := os.OpenFile("aardvarks.txt", options, os.FileMode(0600))
	check(err)
	defer file.Close()

	fmt.Println("read...")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())

	fmt.Println("\nwrite...")
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
