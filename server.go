package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Server running....")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Kubernets!!!</h1>"))
}
