package demos

import (
	"fmt"
)

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
	return sum, sub
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
