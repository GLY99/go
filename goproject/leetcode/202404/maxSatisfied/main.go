package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	// 计算在不生气的情况下顾客数量
	count := 0
	length := len(customers)
	for i := 0; i < length; i++ {
		if grumpy[i] == 0 {
			count += customers[i]
		}
	}

	// 计算如果在最开始的时间不生气，可以增加的用户数
	increase := 0
	for i := 0; i < minutes; i++ {
		increase += customers[i] * grumpy[i]
	}

	// 开始滑动窗口
	maxIncrease := increase
	for i := minutes; i < length; i++ {
		increase = increase - customers[i-minutes]*grumpy[i-minutes] + customers[i]*grumpy[i]
		maxIncrease = Max(increase, maxIncrease)
	}
	return count + maxIncrease
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	fmt.Println(maxSatisfied(customers, grumpy, minutes))
}
