package main

import (
	"errors"
	"fmt"
)

func test1() {
	// defer即使在函数异常后也会被执行
	defer func() {
		// recover()可以捕获程序中的异常并且返回
		err := recover()
		if err != nil {
			fmt.Printf("err=%v\n", err) // err=runtime error: integer divide by zero
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Printf("res=%d\n", res)
}

func test2() {
	// defer即使在函数异常后也会被执行
	defer func() {
		// recover()可以捕获程序中的异常并且返回
		err := recover()
		if err != nil {
			fmt.Printf("err=%v\n", err) // err=panic error
		}
	}()
	// errors.New(string)可以自定义异常
	// panic(any)可以抛出一个异常，这个异常可以被defer recover捕获
	err := errors.New("panic error")
	panic(err)
}

func main() {
	test1()
	fmt.Printf("test1 run finashed!\n")
	test2()
}
