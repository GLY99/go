package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int // 栈最大可以存放的个数
	Top    int // 栈顶
	Arr    []int
}

func (stack *Stack) Push(num int) (err error) {
	if stack.Top+1 == stack.MaxTop {
		err = errors.New("stack full")
		return
	}
	stack.Top++
	stack.Arr = append(stack.Arr, num)
	return
}

func (stack *Stack) Pop() (num int, err error) {
	if stack.Top == -1 {
		num = -1
		err = errors.New("stack empty")
		return
	}
	num = stack.Arr[stack.Top]
	stack.Arr = stack.Arr[:stack.Top]
	stack.Top--
	return
}

func main() {
	stack := Stack{MaxTop: 2, Top: -1}
	fmt.Println(stack.Push(1)) // <nil>
	fmt.Println(stack.Push(2)) // <nil>
	fmt.Println(stack)         // {2 1 [1 2]}
	fmt.Println(stack.Pop())   // 2 <nil>
	fmt.Println(stack.Pop())   // 1 <nil>
	fmt.Println(stack)         // {2 -1 []}
}
