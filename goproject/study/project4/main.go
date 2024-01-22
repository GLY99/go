package main

import "fmt"

func main() {
	// 加号左右都是数值为数学运算
	var num1 int = 1
	var num2 = 2
	num3 := num1 + num2
	fmt.Println("num3:", num3)

	// 加号左右都是字符串为字符串拼接
	var str1 string = "a"
	var str2 = "b"
	str3 := str1 + str2
	fmt.Println("str3:", str3)
}
