package demos

import "fmt"

func Func1() {
	// for 循环基本语法, 这里i的作用域只在for循环中
	for i := 0; i < 10; i++ {
		fmt.Printf("hellow world\n")
	}

	// for循环的写法2, 这里j的作用域在这个函数中
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("hellow world\n")
	}
	fmt.Printf("%d\n", j)
}
