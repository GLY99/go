package main

import "fmt"

func main() {
	// 默认情况下管道可读可写

	// 声明只写chan
	var onlyWriteChan chan<- int = make(chan<- int, 2)
	onlyWriteChan <- 1
	// <-onlyWriteChan // invalid operation: cannot receive from send-only channel onlyWriteChan (variable of type chan<- int)

	// 声明只读chan
	var onlyReadChan <-chan int = make(<-chan int)
	// onlyReadChan <- 1 // invalid operation: cannot send to receive-only channel onlyReadChan (variable of type <-chan int)
	fmt.Println(onlyReadChan)
}
