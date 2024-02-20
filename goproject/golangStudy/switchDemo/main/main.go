package main

import "fmt"

func func1() {
	fmt.Printf("请输入a-g:\n")
	var c byte
	fmt.Scanf("%c\n", &c)
	switch c {
	case 'a':
		fmt.Printf("星期一\n")
	case 'b':
		fmt.Printf("星期二\n")
	case 'c':
		fmt.Printf("星期三\n")
	case 'd':
		fmt.Printf("星期四\n")
	case 'e':
		fmt.Printf("星期五\n")
	case 'f':
		fmt.Printf("星期六\n")
	case 'g':
		fmt.Printf("星期日\n")
	default:
		fmt.Printf("error input\n")
	}
}

func func2() {
	var a int8 = 1
	switch {
	case a == 1:
		fmt.Printf("1\n")
	case a == 2:
		fmt.Printf("2\n")
	}
}

func func3() {
	// switch后面也可以定义声明变量，只能使用:=方式声明
	switch a := 1; {
	case a == 1:
		fmt.Printf("1\n")
	case a == 2:
		fmt.Printf("2\n")
	}
}

func func4() {
	// fallthrough后会继续执行下一个case，不会判断下一个case条件
	switch a := 1; {
	case a == 1:
		fmt.Printf("1\n")
		fallthrough
	case a == 2:
		fmt.Printf("2\n")
	case a == 3:
		fmt.Printf("3\n")
	}
}

func func5() {
	// type switch
	var x interface{}
	var y float32 = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的数据类型是%T\n", i)
	case float32:
		fmt.Printf("x的数据类型是float32\n")
	case int8:
		fmt.Printf("x的数据类型是int8\n")
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	func1()
	func2()
	func3()
	func4()
	func5()
}
