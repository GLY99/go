package main

import (
	"fmt"
	"net/http"
)

// 自定义的handler结构体，需要实现ServeHTTP方法
type MyHandler struct{}

func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my handler")
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	myHandler := MyHandler{}
	// 自定义的多路复用器
	mux := http.NewServeMux()
	// 注册url到不同的handler
	mux.HandleFunc("/helloWorld", handle)
	mux.Handle("/myHandler", &myHandler)
	http.ListenAndServe(":8080", mux)
}
