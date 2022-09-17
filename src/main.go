package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetGreeting)
	http.ListenAndServe(":8080", nil)
}

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*r)
	fmt.Fprintf(w, "Hello")
}
