package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd() failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
	// 解析模板 可以通过Must函数来处理解析模板时的错误
	templateFile1 := path + "/templateWeb1/index1.html"
	templateFile2 := path + "/templateWeb1/index2.html"
	t, err := template.ParseFiles(templateFile1, templateFile2)
	if err != nil {
		fmt.Printf("template.ParseFiles failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
	// 渲染模板2
	err = t.ExecuteTemplate(w, "index2.html", "Hello World")
	if err != nil {
		fmt.Printf("t.ExecuteTemplate failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
}

func main() {
	// 自定义的多路复用器
	mux := http.NewServeMux()

	// 注册url到不同的handler
	mux.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", mux)
}
