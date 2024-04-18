package main

import (
	"errors"
	"fmt"
)

func calculator(num1 int, num2 int, cal string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
	switch cal {
	case "+":
		fmt.Printf("%d+%d=%d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d-%d=%d\n", num1, num2, num1-num2)
	case "*", "x", "X":
		fmt.Printf("%dx%d=%d\n", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			err := errors.New("0不能作为除数")
			panic(err)
		}
		fmt.Printf("%d/%d=%d\n", num1, num2, num1/num2)
	}
}

func main() {
	for {
		fmt.Println("小小计算器")
		fmt.Println("0退出, 1计算")
		var choose int
		fmt.Scanf("%d\n", &choose)
		if choose == 0 {
			break
		}
		var num1, num2 int
		var cal string
		fmt.Printf("请输入类似如下的算式: x+y\n")
		fmt.Scanf("%d %s %d\n", &num1, &cal, &num2)
		calculator(num1, num2, cal)
	}
}
