package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型和string互相转换
func main() {
	// fmt.Sprintf 通过指定格式返回字符串
	// fmt.Sprint 通过默认格式返回字符串,如果两个相邻的参数都不是字符串,会在它们的输出之间添加空格。
	// fmt.Sprintln 通过默认格式返回字符串,总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。
	var num1 int8 = 99
	var num2 float64 = 1.234
	var myChar byte = 'a'
	var myBool bool = true
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("type is %T, str=%s\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("type is %T, str=%s\n", str, str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("type is %T, str=%s\n", str, str)
	str = fmt.Sprintf("%t", myBool)
	fmt.Printf("type is %T, str=%s\n", str, str)

	// 通过strconv
	num3 := int64(num1)
	str = strconv.FormatInt(num3, 10)
	fmt.Printf("type is %T, str=%s\n", str, str)
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("type is %T, str=%s\n", str, str)
	str = strconv.FormatBool(myBool)
	fmt.Printf("type is %T, str=%s\n", str, str)

	//通过strconv的函数将string转成其它基本数据类型
	var str1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("type is %T, b=%t\n", b, b)
	var str2 string = "1234567"
	var num4 int64
	num4, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("type is %T, num4=%d\n", num4, num4)
	var str3 string = "1.23456"
	var num5 float64
	num5, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("type is %T, num5=%f\n", num5, num5)
	var str4 string = "hello"
	var num6 int64
	num6, _ = strconv.ParseInt(str4, 10, 64) // 0
	fmt.Printf("type is %T, num6=%d\n", num6, num6)
}
