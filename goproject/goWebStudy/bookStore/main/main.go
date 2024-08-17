package main

import (
	"fmt"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()
	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		fmt.Printf("http server listen 8080 failed, err=%v\n", err)
	}
}
