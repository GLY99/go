package main

import (
	"fmt"
	"goWebStudy/bookStore/controller"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()
	// 注册路由
	serverMux.HandleFunc("/login", controller.Login)
	serverMux.HandleFunc("/regist", controller.Regsit)

	// 启动监听
	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		fmt.Printf("http server listen 8080 failed, err=%v\n", err)
	}
}
