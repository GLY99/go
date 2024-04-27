package main

import (
	"fmt"
)

func func1() {
	// 如果左右都是整数，go会去掉小数部分保留整数部分。和python整除使用'//'结果y轴向下取整
	fmt.Println(10 / 4)  // 2 如果是python 10 // 4 == 2; 10 // 4.0 == 2.0
	fmt.Println(-10 / 4) // -2 如果是python -10 // 4 == -3; -10 // 4.0 == -3.0
	// 如果希望保留小数部分，go需要有浮点数参与计算。python可以使用'/'除
	fmt.Println(10 / 4.0)  // 2.5 如果是pyton 10 / 4 == 2.5; 10 / 4.0 == 2.5
	fmt.Println(-10 / 4.0) // -2.5 如果是python -10 / 4 == -2.5; -10 / 4.0 == -2.5
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

func func4() {
	// 假设还有97天放假问还有多少个星期多少天
	var days int8 = 97
	var weeks int8 = days / 7
	var day int8 = days % 7
	fmt.Printf("%d星期%d天\n", weeks, day)
}

func func5() {
	// 华氏温度转换摄氏度
	var huaShi float32 = 32
	var sheShi float32 = 5.0 / 9 * (huaShi - 100)
	fmt.Printf("%v对应的摄氏温度是%v\n", huaShi, sheShi)
}

func func6() {
	// 关系运算符案例
	var num1 int8 = 1
	var num2 int8 = 2
	flag := num1 < num2
	fmt.Println(flag) // true
	flag = num1 < num2
	fmt.Println(flag) // true
	flag = num1 == num2
	fmt.Println(flag) // false
	flag = num1 != num2
	fmt.Println(flag) // true
	flag = num1 <= num2
	fmt.Println(flag) // true
}

func func7() {
	var age int = 32
	if age >= 18 && age <= 60 {
		fmt.Println("ok1")
	}
	if !(age > 60) {
		fmt.Println("ok2")
	}
	if age > 18 || age < 0 {
		fmt.Println("ok3")
	}
}

func func8() {
	num1 := 1
	num2 := 2
	var tempNum int = num1
	num1 = num2
	num2 = tempNum
	fmt.Println(num1, num2)
}

func func9() {
	num := 100
	var ptr *int = &num
	var ptrToPtr **int = &ptr
	fmt.Println("num的地址即ptr的值是:", ptr)
	fmt.Println("ptr的地址是:", &ptr)
	fmt.Println("prt指针对应的地址中存储的值是:", *ptr)
	fmt.Println("ptrToPtr的值是:", ptrToPtr)
	fmt.Println("ptrToPtr的地址是:", &ptrToPtr)
	fmt.Println("ptrToPtr指针对应的地址存储的值是:", *ptrToPtr)
}

func main() {
	func1()
	func2()
	func3()
	func4()
	func5()
	func6()
	func7()
	func8()
	func9()
}
