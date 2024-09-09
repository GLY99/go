package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func testIf(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd() failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
	// 解析模板 可以通过Must函数来处理解析模板时的错误
	templateFile := path + "/templateWeb2/if.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Printf("template.ParseFiles failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
	// 渲染模板2
	err = t.ExecuteTemplate(w, "if.html", false)
	if err != nil {
		fmt.Printf("t.ExecuteTemplate failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
}

func testRange(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd() failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
	// 解析模板 可以通过Must函数来处理解析模板时的错误
	templateFile := path + "/templateWeb2/range.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Printf("template.ParseFiles failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
	// 渲染模板2
	var slice []int
	slice = []int{1, 2, 3}
	err = t.ExecuteTemplate(w, "range.html", slice)
	if err != nil {
		fmt.Printf("t.ExecuteTemplate failed, err=%v\n", err)
		w.Write([]byte("Server error"))
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test_if", testIf)
	mux.HandleFunc("/test_range", testRange)
	http.ListenAndServe(":8080", mux)
}
