package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my handler")
}

func main() {
	myHandler := MyHandler{}

	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("listen server failed, err=%v", err)
	}
}
