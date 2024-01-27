package main

import (
	"fmt"
)

// str使用
func main() {
	// 字符串是不可变的
	var s1 string = "hello world"
	fmt.Println(s1)

	// 字符串中有特殊字符，可以使用反引号将字符串原生输出
	s1 = `hello\nworld`
	fmt.Println(s1)
	s2 := "hello\nworld"
	fmt.Println(s2)

	// 使用+进行字符串拼接, 需要将+保留在上面
	s3 := "hello" +
		"world"
	fmt.Println(s3)
}
