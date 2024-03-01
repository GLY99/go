package demos

import (
	"fmt"
)

func init() {
	fmt.Printf("init funcDemos\n")
}

func Cal(num1 float64, num2 float64, operator byte) (res float64) {
	// var res float64  // 如果在函数的返回类型列表中指定了变量相当于定义了一个变量，函数内不需要重复定义
	switch operator {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		res = num1 / num2
	default:
		fmt.Printf("error operator...\n")
	}
	return res
}

func GetSumAndSub(num1 float64, num2 float64) (sum float64, sub float64) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func FeboNums(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return FeboNums(num-1) + FeboNums(num-2)
}

func GetFuncAsn(num int) int {
	if num == 1 {
		return 3
	}
	return 2*GetFuncAsn(num-1) + 1
}

func GetPeachs(day int) int {
	if day > 10 || day < 1 {
		return -1
	}
	if day == 10 {
		return 1
	}
	return 2 * (GetPeachs(day+1) + 1)
}

func TestValueTransfer(num int) {
	num++
}

func TestQuoteTransfer(num *int) {
	*num++
}

func GetSum(num1 int, num2 int) (res int) {
	res = num1 + num2
	return res
}

func TestFuncTransfer1(myFunc func(num1 int, num2 int) (res int), num1 int, num2 int) int {
	return myFunc(num1, num2)
}

func TestFuncTransfer2(myFunc func(int, int) int, num1 int, num2 int) int {
	return myFunc(num1, num2)
}

type myFuncType func(int, int) int

func TestFuncTransfer3(myFunc myFuncType, num1 int, num2 int) int {
	return myFunc(num1, num2)
}

func GetManyNumsSum1(args ...int) int {
	var sum int
	for _, val := range args {
		sum = sum + val
	}
	return sum
}

func GetManyNumsSum2(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum = sum + args[i]
	}
	return sum
}

func Swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}
