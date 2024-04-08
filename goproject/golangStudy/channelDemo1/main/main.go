package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// channel make后才能使用
	var channel1 chan int = make(chan int, 3)
	defer close(channel1)
	fmt.Printf("channel1的值%v, channel1本身的地址%v\n", channel1, &channel1) // channel1的值0xc00010c000, channel1本身的地址0xc00006c020

	// 向channel写入数据
	channel1 <- 1
	num1 := 2
	channel1 <- num1

	// 管道不会自动扩容，所以存放数据不能超过cap, 否则报错
	channel1 <- 3
	// channel1 <- 4 // fatal error: all goroutines are asleep - deadlock!

	// 查看chan size和cap
	fmt.Printf("channel1 size=%d, cap=%d\n", len(channel1), cap(channel1)) // channel1 size=3, cap=3

	// 从channel取数据
	num2 := <-channel1
	fmt.Println(num2)                                                      // 1
	fmt.Printf("channel1 size=%d, cap=%d\n", len(channel1), cap(channel1)) // channel1 size=2, cap=3

	// 在没有使用协程的情况下，如果管道数据全部取出，再取就会报错
	num3 := <-channel1
	num4 := <-channel1
	// num5 := <-channel1      // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(num3, num4) // 2 3

	// 接口类型的chan
	chanEveryType := make(chan interface{}, 5)
	chanEveryType <- 10
	chanEveryType <- "abc"
	chanEveryType <- &Cat{Name: "mimi", Age: 10}

	<-chanEveryType
	<-chanEveryType
	cat := <-chanEveryType
	newCat, ok := cat.(*Cat)
	if !ok {
		fmt.Printf("cat is not *Cat type\n")
	} else {
		fmt.Println(newCat.Name)
	}
}
