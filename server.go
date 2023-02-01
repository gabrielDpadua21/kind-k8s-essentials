package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Print("Server running....")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	color := os.Getenv("COLOR")
	fmt.Fprintf(w, "Hello, %s. The kubernets has you, follow the %s habbit", name, color)
}
