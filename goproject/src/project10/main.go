package main

import (
	"fmt"
)

// 基本数据类型默认值 基本数据类型转换
func main() {
	var num1 int     // 0
	var str string   // ""
	var num2 float32 // 0
	var num3 float64 // 0
	var b bool       // false
	fmt.Println(num1, num2, num3, str, b)

	// 必须显示转换
	var num4 int = 64
	var num5 float32 = float32(num4)
	var num6 uint = uint(num4)
	num7 := int32(num4)
	fmt.Println(num4, num5, num6, num7)

	// 数值超出转换后的类型范围会溢出,溢出后二进制截断
	var num8 int32 = 999
	var num9 int8 = int8(num8)
	fmt.Println(num8, num9)
}
