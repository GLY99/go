package main

import "fmt"

func sum(n1, n2 int) int {
	// 当执行到defer时，会先将后面的执行语句压入一个单独的栈中（defer栈）
	// 当前函数执行完毕时，从defer栈中按照先入后出的方式执行语句
	defer fmt.Printf("n1=%d\n", n1) // n1 == 1
	defer fmt.Printf("n2=%d\n", n2)

	res := n1 + n2
	n1++
	fmt.Printf("func sum() res=%d, n1=%d\n", res, n1) // 3, n1 == 2
	return res
}

func main() {
	sum(1, 2)
	fmt.Printf("finash!")
	// func sum() res=3, n1=2
	// n2=2
	// n1=1
	// finash!
}
