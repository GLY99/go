package stack

import (
	"errors"
)

type Stack struct {
	MaxTop int // 栈最大可以存放的个数
	Top    int // 栈顶
	Arr    []int
}

func NewStack(size int) *Stack {
	return &Stack{MaxTop: size, Top: -1}
}

func (stack *Stack) Push(num int) (err error) {
	if stack.IsFull() {
		err = errors.New("stack full")
		return
	}
	stack.Top++
	stack.Arr = append(stack.Arr, num)
	return
}

func (stack *Stack) Pop() (num int, err error) {
	if stack.IsEmpty() {
		num = -1
		err = errors.New("stack empty")
		return
	}
	num = stack.Arr[stack.Top]
	stack.Arr = stack.Arr[:stack.Top]
	stack.Top--
	return
}

func (stack *Stack) IsFull() bool {
	return stack.Top+1 == stack.MaxTop
}

func (stack *Stack) IsEmpty() bool {
	return stack.Top == -1
}
