package main

import "fmt"

// 字符使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	// go的字符存储在byte类型中，直接输出，会输出其对应的ASCII码值
	fmt.Println("c1 is:", c1, "c2 is:", c2)
	// 格式化输出会输出字符值
	fmt.Printf("c1 is:%c, c2 is:%c\n", c1, c2)
	var c3 int16 = '北'
	fmt.Println("c3 is:", c3)
	fmt.Printf("c3 is:%c", c3)
}
