package main

import (
	"fmt"
	"math/rand"
)

func guessNumber() func(int, int) (int, bool) {
	var count int = 0
	return func(num int, expectNum int) (int, bool) {
		result := false
		if num > expectNum {
			fmt.Println("猜大了！")
		} else if num < expectNum {
			fmt.Println("猜小了！")
		} else {
			fmt.Println("猜对了！")
			result = true
		}
		count++
		return count, result
	}
}

func main() {
	num := rand.Intn(100)
	myFunc := guessNumber()
	for i := 0; i < 10; i++ {
		var inputNum int
		fmt.Printf("请输入猜测的数字\n")
		fmt.Scanln(&inputNum)
		count, result := myFunc(inputNum, num)
		if result {
			if count == 1 {
				fmt.Printf("天才！\n")
			} else if count >= 2 && count <= 3 {
				fmt.Printf("你很聪明！\n")
			} else if count >= 4 && count <= 9 {
				fmt.Printf("你一般！\n")
			} else {
				fmt.Printf("终于猜对了！\n")
			}
			break
		} else {
			if count == 10 {
				fmt.Printf("说你点啥呢！\n")
			}
		}
	}
}
