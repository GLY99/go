package main

import (
	"fmt"
)

func func1() {
	// 通过指针修改变量的值
	var num int = 10
	var ptr *int = &num
	*ptr = 20
	fmt.Print(num, *ptr)
}

func func2() {
	a := 1
	b := &a            // 取a变量地址给b,a b最终都指向1的内存空间
	*b = 2             // a b直接会互相影响
	fmt.Println(a, *b) // 2 2
	a = 3
	fmt.Println(a, *b)    // 3 3
	c := *b               // 取b指针指向的内存地址的值给c,这里相当于值copy
	c = 4                 // c修改不会影响a b
	fmt.Println(a, *b, c) //3 3 4
}

// 基本数据类型和string互相转换
func main() {
	// 通过&获取变量的地址
	var num int = 10
	fmt.Println("num的地址=", &num)

	// ptr是一个指针变量，类型*int
	// 指针变量存储的是地址
	var ptr *int = &num
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Println("ptr的地址=", &ptr)

	// 使用*获取指定变量指向的值
	fmt.Printf("ptr指向的地址存储的值是%v\n", *ptr)

	func1()
	func2()
}
