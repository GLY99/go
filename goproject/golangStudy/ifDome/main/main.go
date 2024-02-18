package main

import (
	"fmt"
)

func func1() {
	var age int8
	fmt.Printf("请输入年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Printf("你的年龄大于18\n")
	} else {
		fmt.Printf("你的年龄未满18岁\n")
	}
}

func func2() {
	var num1 int32 = 10
	var num2 int32 = 30
	if num1+num2 > 20 {
		fmt.Printf("hello world\n")
	}
}

func func3() {
	var num1 float32 = 12.5
	var num2 float32 = 13.5
	if num1 > 10 && num2 < 20 {
		fmt.Printf("hello world\n")
	}
}

func func4() {
	var num1 int32 = 10
	var num2 int32 = 30
	if (num1+num2)%3 == 0 && (num1+num2)%5 == 0 {
		fmt.Printf("yes\n")
	} else {
		fmt.Printf("no\n")
	}
}

func func5() {
	var year int32 = 2014
	if year%4 == 0 && year%100 != 0 && year%400 == 0 {
		fmt.Printf("run nian\n")
	} else {
		fmt.Printf("no run nian\n")
	}
}

func main() {
	func1()
	func2()
	func3()
	func4()
	func5()
}
