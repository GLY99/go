package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my handler")
}

func main() {
	http.Handle("/myHandler", &MyHandler{})
	http.ListenAndServe(":8080", nil)
}
