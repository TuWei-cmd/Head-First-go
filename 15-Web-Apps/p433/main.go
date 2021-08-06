package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, Web")
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, Web")
}
func chineseHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "你好, Web")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/Salut", frenchHandler)
	http.HandleFunc("/Namaste", hindiHandler)
	http.HandleFunc("/zh", chineseHandler)

	err := http.ListenAndServe("localhost:8888", nil)
	log.Fatal(err)
}
