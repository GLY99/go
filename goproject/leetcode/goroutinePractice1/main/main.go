package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	defer close(intChan)
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("write data %v\n", i)
		intChan <- i
	}
}

func readData(intChan chan int, flagChan chan bool) {
	defer close(flagChan)
	for {
		v, ok := <-intChan // 管道关闭时，读取完所有元素再次读取ok=false, 协程的情况下读取空管道会阻塞
		if !ok {
			break
		}
		fmt.Printf("read data %v\n", v)
	}
	flagChan <- true
}

func main() {
	intChan := make(chan int, 50)
	flagChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, flagChan)

	for {
		v, ok := <-flagChan // 协程的情况下读取空管道会阻塞
		if v || !ok {
			break
		}
	}
}
