package main

import (
	"fmt"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()
	// 处理css js 图片等静态资源
	// http.StripPrefix 将所有以/static开头的url去除这个前缀后交给后面的handle处理
	// http.FileServer 返回一个查找文件的handle, 将会在http.Dir指定的路径中查找对应的资源
	// http.Dir 返回一个http.File类型的变量交给http.FileServer使用
	serverMux.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("/home"))))
	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		fmt.Printf("http server listen 8080 failed, err=%v\n", err)
	}
}
