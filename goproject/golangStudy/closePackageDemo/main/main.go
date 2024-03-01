package main

import (
	"fmt"
	"golangStudy/closePackageDemo/demos"
)

func addUpper() func(int) int {
	// addUpper函数返回一个匿名函数
	// 这个匿名函数引用了函数外的变量n
	// 因此这个匿名函数和变量n组成一个整体，形成闭包
	// 可以理解为闭包是一个类，匿名函数是一个类的方法，引用的变量是类的变量
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	// 在调用addUpper时n完成初始化，返回一个匿名函数
	// 后面每次调用这个匿名函数n的值都会被更新
	f1 := addUpper()          // 将func(x int) int 函数赋值给f
	fmt.Printf("%d\n", f1(1)) // 11
	fmt.Printf("%d\n", f1(2)) // 13
	fmt.Printf("%d\n", f1(3)) // 16

	// 闭包实践
	f2 := demos.MakeSuffix(".pcap")
	fmt.Printf("%s\n", f2("a.pcap"))
	fmt.Printf("%s\n", f2("a.jpg"))
}
