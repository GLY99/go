package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		intChan <- i
	}
}

func readData(intChan <-chan int, exitChan chan<- bool) {
	defer func() {
		exitChan <- true
	}()
	for {
		select {
		case v := <-intChan: // 这里如果读取不到数据直接跳过，执行default
			fmt.Printf("read data %d from intChan\n", v)
		default:
			fmt.Printf("no data read!\n") // intChan里面的数据都读取完了这里会一直被执行，不会deadlock
		}
	}
}

func main() {
	intChan := make(chan int, 5)
	exitChan := make(chan bool, 1)
	defer close(exitChan)
	defer close(intChan)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// 下面两种阻塞方式二选一都可以
	// <-exitChan // 阻塞
	for {
		select {
		case <-exitChan: // 没有数据可读也不阻塞
			return
		}
	}
}
