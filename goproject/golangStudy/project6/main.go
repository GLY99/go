package main

import "fmt"

// 小数使用
func main() {
	var price float32 = 32.23
	fmt.Println("price is:", price)

	var num1 float32 = -0.00032
	var num2 float64 = -78909322.26
	fmt.Println("num1 is:", num1, ",num2 is:", num2)

	// 浮点数可能会造成尾数部分丢失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3 is:", num3, ",num4 is:", num4)

	// float默认64,并且不受操作系统影响
	var num5 = -123.0000901
	fmt.Printf("num5 type: %T\n", num5)

	// 十进制数形式
	num6 := 5.12
	num7 := .512
	fmt.Println("num6 is:", num6, ",num7 is:", num7)

	// 科学计数法
	num8 := 5.1234e2   // => 5.1234 * 10^2
	num9 := 5.1234e2   // => 5.1234 * 10^2
	num10 := 5.1234e-2 // => 5.1234 / 10^2
	fmt.Println("num8 is:", num8, ",num9 is:", num9, "num10 is:", num10)
}
