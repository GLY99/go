package main

import (
	"fmt"
	"reflect"
)

func test1(n1, n2 int) {
	fmt.Printf("test1 %d, %d\n", n1, n2)
}

func test2(n1, n2 int, str string) {
	fmt.Printf("test2 %d %d %s\n", n1, n2, str)
}

// 通过反射实现函数适配器
func BridgeFunc(call interface{}, args ...interface{}) {
	// 获取传入的参数数量并且将其转化成Value类型
	n := len(args)
	funcParams := make([]reflect.Value, n)
	for i := 0; i < n; i++ {
		funcParams[i] = reflect.ValueOf(args[i])
	}
	// 获取函数，然后调用
	function := reflect.ValueOf(call)
	function.Call(funcParams)
}

func main() {
	BridgeFunc(test1, 1, 2)                   // test1 1, 2
	BridgeFunc(test2, 3, 4, "this is string") // test2 3 4 this is string
}
