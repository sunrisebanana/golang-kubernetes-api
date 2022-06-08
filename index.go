package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello User!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", index)

	fmt.Print("Starting...")
	http.ListenAndServe(":8080", nil)
	fmt.Print("Started")
}
