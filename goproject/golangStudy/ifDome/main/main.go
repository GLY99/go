package main

import (
	"fmt"
)

func func1() {
	var age int8
	fmt.Printf("请输入年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Printf("你的年龄大于18")
	} else {
		fmt.Printf("你的年龄未满18岁")
	}
}

func main() {
	func1()
}
