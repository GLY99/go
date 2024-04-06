package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func PrintHelloWorld(wg *sync.WaitGroup) {
	defer wg.Done() // 通知一个协程任务结束了
	for i := 0; i < 10; i++ {
		fmt.Println("PrintHelloWorld(): hello world" + strconv.Itoa(i+1))
		time.Sleep(2 * time.Second)
	}
}

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Printf("cpu num=%d\n", cpuNum)    // 获取当前机器的逻辑cpu个数
	res := runtime.GOMAXPROCS(cpuNum / 2) // 可以自己手动设置最多使用多少个cpu
	fmt.Printf("set max procs=%d\n", res)

	wg := sync.WaitGroup{}
	wg.Add(1)               // 标记启动了一个协程
	go PrintHelloWorld(&wg) // 在主线程中开启一个协程, 如果主线程结束,协程也会结束
	for i := 0; i < 10; i++ {
		fmt.Println("main(): hello world" + strconv.Itoa(i+1))
		time.Sleep(1 * time.Second)
	}
	wg.Wait() // 当所有协程的任务都结束这里才会继续往下执行,否则阻塞。
}
