package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hello!"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)
}
