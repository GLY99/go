package main

import (
	"fmt"
	demo "golangStudy/funcDemo/demos" // demo是别名
)

func main() {
	res := demo.Cal(1, 2, '+')
	fmt.Printf("%.2f\n", res)

	sum, _ := demo.GetSumAndSub(1.2, 0.2)
	_, sub := demo.GetSumAndSub(1.2, 0.2)
	fmt.Printf("sum=%.2f, sub=%.2f\n", sum, sub)

	num := demo.FeboNums(7)
	fmt.Printf("febonums=%d\n", num)

	num = demo.GetFuncAsn(5)
	fmt.Printf("getfuncasn=%d\n", num)

	num = demo.GetPeachs(1)
	fmt.Printf("getpeachs=%d\n", num)

	n := 1
	demo.TestValueTransfer(n)
	fmt.Printf("%d\n", n) // 1
	demo.TestQuoteTransfer(&n)
	fmt.Printf("%d\n", n) // 2

	// 函数也是一种数据类型，可以赋值被变量
	myFunc := demo.Cal
	fmt.Printf("myFunc的数据类型是%T, demo.Cal的数据类型是%T\n", myFunc, demo.Cal)
	fmt.Printf("%2.f\n", myFunc(2, 2, '-'))

	// 函数也可以传递给另外一个函数，类型以func作为关键字，后面加上型参类型和返回值类型，形参变量和返回值变量可选
	fmt.Printf("%d\n", demo.TestFuncTransfer1(demo.GetSum, 1, 1))
	fmt.Printf("%d\n", demo.TestFuncTransfer2(demo.GetSum, 1, 1))
	fmt.Printf("%d\n", demo.TestFuncTransfer3(demo.GetSum, 1, 1))

	// type的使用
	type myInt int
	var num1 myInt = 10
	var num2 int = int(num1)           // 虽然使用效果上这两个类型没有区别，但是它们依然时两个不同类型
	fmt.Printf("%d, %d\n", num1, num2) // 10, 10

	fmt.Printf("%d\n", demo.GetManyNumsSum1(1, 2, 3)) // 6
	fmt.Printf("%d\n", demo.GetManyNumsSum2(1, 2, 3)) // 6
}
