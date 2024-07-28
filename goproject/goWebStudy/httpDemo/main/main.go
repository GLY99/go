package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "httpDemo")
}

func main() {
	http.HandleFunc("/http", handle)
	http.ListenAndServe(":8080", nil)
}
