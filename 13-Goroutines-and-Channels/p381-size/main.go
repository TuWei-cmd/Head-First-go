package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Size of ", url, ":", len(body))
}

func main() {
	go responseSize("https://golang.google.cn/doc/")
	go responseSize("http://example.com/")
	go responseSize("https://golang.google.cn/")
	go responseSize("https://cn.bing.com/")
	time.Sleep(5 * time.Second)
}
