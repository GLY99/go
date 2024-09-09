package main

import (
	"fmt"
	"goWebStudy/bookStore/route"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()

	// 注册路由
	route.RegistRoute(serverMux)

	// 启动监听
	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		fmt.Printf("http server listen 8080 failed, err=%v\n", err)
	}
}
