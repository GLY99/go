package main

import (
	"fmt"
)

func main() {
	// chan关闭后只能读不能写
	intChan := make(chan int, 5)
	intChan <- 1
	intChan <- 2
	close(intChan)
	// intChan <- 3 // panic: send on closed channel
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println(num1, num2, num3) // 1 2 0, 0是一个默认值

	// 管道遍历，管道没有close遍历会报错，关闭后可以正常读数据
	intChan2 := make(chan int, 100)
	for i := 0; i < cap(intChan2); i++ {
		intChan2 <- i
	}
	close(intChan2) // 如果没有这个下面遍历会报 fatal error: all goroutines are asleep - deadlock!
	for v := range intChan2 {
		fmt.Println(v)
	}
}
