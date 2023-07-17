package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Server!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting HTTP Server")
	http.ListenAndServe(":9091", nil)
}
