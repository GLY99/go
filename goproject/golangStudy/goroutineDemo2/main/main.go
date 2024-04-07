package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int, 20)
	// 声明一个全局的互斥锁
	lock = &sync.Mutex{}
)

func test(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 写的时候加锁
	lock.Lock()
	myMap[n] = res // fatal error: concurrent map writes
	// 写完释放锁
	lock.Unlock()
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go test(wg, i)
	}
	wg.Wait()
	for k, v := range myMap {
		fmt.Printf("%d!=%d\n", k, v)
	}
}
