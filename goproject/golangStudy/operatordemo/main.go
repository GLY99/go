package main

import "fmt"

func func1() {
	// 如果左右都是整数，结果向下取整，即去掉小数部分保留整数部分
	fmt.Println(10 / 4) // 2
	// 如果希望保留小数部分，需要有浮点数参与计算
	fmt.Println(10 / 4.0) // 2.5
}

func func2() {
	// 取模公式 a % b = a - a / b * b
	// 模的结果符号总是和被取模数的符号一致，10 % -3 10是被取模数
	fmt.Println(10 % 3)   // 1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(-10 % -3) // -1
}

func func3() {
	// golang ++ -- 只能当作一个单独语句使用
	var num int8 = 1
	num++
	fmt.Println(num) // 2
	num--
	fmt.Println(num) // 1
	// tempNum := num++ // 错误用法
}

func main() {
	func1()
	func2()
	func3()
}
