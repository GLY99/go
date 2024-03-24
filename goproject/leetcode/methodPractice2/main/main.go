package main

import (
	"errors"
	"fmt"
	"strconv"
)

type MethodUtil struct {
	// 空结构体
}

type Calcautor struct {
	Num1 float64
	Num2 float64
}

func (c *Calcautor) CalcautorNum(operator byte) (float64, error) {
	var res float64
	var err error
	switch operator {
	case '+':
		res = c.Num1 + c.Num2
	case '-':
		res = c.Num1 - c.Num2
	case '*':
		res = c.Num1 * c.Num2
	case '/':
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", c.Num1/c.Num2), 64)
	default:
		err = errors.New("unkonwn operator")
	}
	return res, err
}

func (mu *MethodUtil) Print1() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtil) Print2(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtil) Print3(m, n int, c string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s", c)
		}
		fmt.Println()
	}
}

func (mu *MethodUtil) Area(len, width float64) float64 {
	return len * width
}

func (mu *MethodUtil) JudgeNum(num int) bool {
	return num%2 == 0
}

func main() {
	mu := MethodUtil{}
	mu.Print1()
	mu.Print2(2, 2)
	mu.Print3(2, 2, "+")
	fmt.Println(mu.Area(2, 2))
	fmt.Println(mu.JudgeNum(1))

	cal := Calcautor{1, 2}
	var operators [5]byte = [5]byte{'+', '-', '*', '/', '?'}
	for _, operator := range operators {
		fmt.Println(cal.CalcautorNum(operator))
	}
}
