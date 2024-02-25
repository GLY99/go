package demos

import "fmt"

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
