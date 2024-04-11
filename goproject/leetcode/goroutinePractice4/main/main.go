package main

import (
	"fmt"
	"runtime"
)

func writeNum(intChan chan int) {
	defer close(intChan)
	fmt.Printf("start write num to intChan\n")
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	fmt.Printf("finash write num to intChan\n")
}

func isPrimeNum(num int) bool {
	if num == 1 {
		return true
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func collectNum(intChan chan int, primeChan chan int, flagChan chan bool) {
	defer func() {
		fmt.Printf("有一个collectNum协程退出\n")
		flagChan <- true
	}()

	var v int
	var ok bool
	for {
		v, ok = <-intChan
		if !ok {
			break
		}
		if isPrimeNum(v) {
			primeChan <- v
		}
	}
}

func waitCollectFinash(flagChan chan bool, primeChan chan int, n int) {
	defer close(primeChan)
	defer close(flagChan)
	for i := 0; i < n; i++ {
		<-flagChan
	}
}

func main() {
	cpuNum := runtime.NumCPU()                   // 获取当前机器的逻辑cpu个数
	usedCpuNum := runtime.GOMAXPROCS(cpuNum / 2) // 可以自己手动设置最多使用多少个cpu
	fmt.Printf("used cpu num=%d\n", usedCpuNum)
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	flagChan := make(chan bool, usedCpuNum)
	go writeNum(intChan)
	for i := 0; i < usedCpuNum; i++ {
		go collectNum(intChan, primeChan, flagChan)
	}
	go waitCollectFinash(flagChan, primeChan, usedCpuNum)
	count := 0
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(v, ok)
		count++
	}
	fmt.Printf("1-8000共有%d个素数\n", count) // 1-8000共有1008个素数
}
