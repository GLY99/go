package main

import "fmt"

type MyError struct {
	errInfo string
}

// 在go中error是接口类型，所以自定义错误类型只需要实现error接口中所有方法就可以将自定义错误类型传递给error类型
// type error interface {
//	Error() string
// }

func (err *MyError) Error() string {
	return err.errInfo
}

// 抛出自定义错误类型并且捕获
func test1() {
	defer func() {
		// recover()可以捕获程序中的异常并且返回
		err := recover()
		if err != nil {
			fmt.Printf("err=%v\n", err) // err=自定义错误类型2
		}
	}()
	err := &MyError{"自定义错误类型2"}
	panic(err) // 抛出的类型要和Error()方法绑定的类型一致，不然显示可能有点奇怪
}

// 返回自定义错误类型
func test2() error {
	err := &MyError{"自定义错误类型3"}
	return err // 将自定义错误类型返回给error接口
}

func main() {
	myerr := &MyError{"自定义错误类型1"}
	var err error
	err = myerr             // 因为是*MyError实现了接口的方法，所以传递的值是地址
	fmt.Printf("%v\n", err) // 自定义错误类型1
	test1()
	err = test2()
	fmt.Printf("%v\n", err) // 自定义错误类型3
}
