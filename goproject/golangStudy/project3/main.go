package main

import "fmt"

// 定义全局变量, 不能用:=
var global_num1 int = 100
var global_num2 = 200

var (
	gloabl_num3 int = 300
	gloabl_num4     = 400
)

func main() {
	fmt.Println("gloabl num is:", global_num1, global_num2, gloabl_num3, gloabl_num4)

	// 声明且赋值
	var num1 int = 1
	fmt.Println("num1 is:", num1)

	// 先声明后赋值
	var num2 int
	num2 = 2
	fmt.Println("num2 is:", num2)

	// 类型推导
	var num3 = "tom"
	fmt.Println("num3 is:", num3)

	// 省略var 使用:=
	// 等价于var num4 int; num4 = 4 因此这个变量不能在别的地方被声明，避免多次声明
	num4 := 4
	fmt.Println("num4 is:", num4)

	// 多变量声明
	var num5, num6, num7 int
	fmt.Println("num5, num6, num7:", num5, num6, num7)
	var num8, num9 = 8, "9"
	fmt.Println("num8, num9:", num8, num9)
	num10, num11 := 10, 11
	fmt.Println("num10, num11:", num10, num11)
}
