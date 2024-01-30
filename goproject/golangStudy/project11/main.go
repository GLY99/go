package main

import (
	"fmt"
)

func func1() {
	var n1 int8 = 100
	var n2 int64
	var n3 int32
	n2 = int64(n1) + 20
	n3 = int32(n1) + 20
	fmt.Println(n1, n2, n3)
}

func func2() {
	var n1 int32 = 666
	var n2 int8
	var n3 int8
	n2 = int8(n1) + 20
	n3 = int8(n1) + 20 // -82
	fmt.Println(n1, n2, n3)
}

// 基本数据类型默认值 基本数据类型转换
func main() {
	func1()
	func2()
}
