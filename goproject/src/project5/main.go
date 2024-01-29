package main

import (
	"fmt"
	"unsafe"
)

// 整数使用
func main() {
	// int 有符号 32位系统4字节，64位系统8字节
	// uint 无符号 32位系统4字节，64位系统8字节
	// rune 有符号 等价int32表示一个unicode码
	// byte 无符号 等价uint8
	var num1 int = 1
	fmt.Println("num1:", num1)
	var num_byte byte = 'a'
	fmt.Println("num_byte:", num_byte)

	// int8 -128-127,int16 int32 int64类推
	var num2 int8 = -128
	var num3 int8 = 127
	fmt.Println("int8 max is", num3, "min is", num2)

	// uint8 0-255,uint16 uint32 uint64类推
	var num4 uint8 = 0
	var num5 uint8 = 255
	fmt.Println("uint8 max is", num4, "min is", num5)

	// 默认是int
	var num6 = 100
	fmt.Printf("num6 is %d\n", num6)
	fmt.Printf("num6 type is %T\n", num6)

	// 查看一个数占用的字节大小
	var num7 int64 = 200
	fmt.Printf("num7 used %d bytes\n", unsafe.Sizeof(num7))
}
