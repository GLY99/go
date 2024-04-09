package main

import "fmt"

func writeData(numChan chan int) {
	defer close(numChan)
	for i := 1; i <= 2000; i++ {
		fmt.Printf("Write data %d\n", i)
		numChan <- i
	}
}

func readData(numChan chan int, resChan chan int, closeResChanFlag chan bool) {
	// 因为有8个协程会调用这个函数，需要保证所有协程都写完数据后关闭resChan 否则会写入失败
	for {
		v, ok := <-numChan
		fmt.Printf("Read data %d, %v\n", v, ok)
		if !ok {
			break
		}
		res := 0
		for i := 1; i <= v; i++ {
			res += i
		}
		resChan <- res
	}
	closeResChanFlag <- true
}

func readResChan(resChan chan int, flagChan chan bool) {
	defer close(flagChan)
	i := 1
	for v := range resChan {
		fmt.Printf("res[%d]=%d\n", i, v)
		i++
	}
	flagChan <- true
}

func closeResChan(resChan chan int, closeResChanFlag chan bool) {
	defer close(closeResChanFlag)
	defer close(resChan)
	count := 0
	for {
		if count == 8 {
			break
		}
		v := <-closeResChanFlag
		if v {
			count++
		}
	}
}

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan int, 2000)
	closeResChanFlag := make(chan bool, 8)
	flagChan := make(chan bool, 1)
	go writeData(numChan)
	for i := 0; i < 8; i++ {
		go readData(numChan, resChan, closeResChanFlag)
	}
	go readResChan(resChan, flagChan)
	go closeResChan(resChan, closeResChanFlag)
	for {
		v, ok := <-flagChan
		if v || !ok {
			break
		}
	}
}
