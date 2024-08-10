package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 自定义的handler结构体，需要实现ServeHTTP方法
type MyHandler struct{}

type RequestBody struct {
	Key3 string
	Key4 string
}

func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 假设请求是POST http://localhost:8080/hello?key1=hello1&key2=hello2 {
	//     "key3": "test",
	//     "key4": "test"
	// }

	// 获取uri
	path := r.URL.Path // / hello
	fmt.Println(path)

	// 获取params
	params := r.URL.RawQuery // key1=hello1&key2=hello2
	fmt.Println(params)

	// 快速获取params, 需要在ParseForm后才能使用Form
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("ParseForm failed, err=%s", err)
	} else {
		params2 := r.Form
		fmt.Println(params2) // map[key1:[hello1] key2:[hello2]]
	}

	// 直接获取params中的值, 不使用ParseForm
	fmt.Println(r.FormValue("key1")) // hello1

	// 获取headers
	headers := r.Header
	// map[
	// Accept:[*/*]
	// Accept-Encoding:[gzip, deflate, br]
	// Connection:[keep-alive]
	// Content-Length:[0]
	// Postman-Token:[9d5385c9-afb1-4005-a43a-d50abbc1d69e]
	// User-Agent:[PostmanRuntime/7.40.0]
	// ]
	fmt.Println(headers)

	// 获取request body
	length := r.ContentLength
	byteSlice := make([]byte, length)
	r.Body.Read(byteSlice)
	rb := &RequestBody{}
	json.Unmarshal(byteSlice, rb)
	fmt.Println(rb) // &{test test}

	fmt.Fprintf(w, "path: %s, params: %s, headers: %s, request body: %s", path, params, headers, rb)
}

func main() {
	myHandler := MyHandler{}

	// 自定义的多路复用器
	mux := http.NewServeMux()

	// 注册url到不同的handler
	mux.Handle("/hello", &myHandler)
	http.ListenAndServe(":8080", mux)
}
