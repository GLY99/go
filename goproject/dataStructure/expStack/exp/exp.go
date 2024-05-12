package exp

import (
	"dataStructure/expStack/stack"
	"errors"
	"strconv"
)

func Exp(str string) (int, error) {
	length := len(str)
	numStack := stack.NewStack(length)
	operStack := stack.NewStack(length)
	idx := 0
	// 字符串由字节组成，字符串中存储的是字节对应的ASCII码值
	for ; idx < length; idx++ {
		tempVal := int(str[idx]) // 字符对应的ASCII码值
		if isOper(tempVal) {
			if operStack.IsEmpty() {
				operStack.Push(tempVal)
			} else {
				curPriority, err1 := priority(tempVal)
				topPriority, err2 := priority(operStack.Arr[operStack.Top])
				if err1 != nil || err2 != nil {
					return -1, errors.Join(err1, err2)
				}
				if topPriority >= curPriority {
					num2, err1 := numStack.Pop()
					num1, err2 := numStack.Pop()
					oper, err3 := operStack.Pop()
					if err1 != nil || err2 != nil || err3 != nil {
						return -1, errors.Join(err1, err2, err3)
					}
					res, err := cal(num1, num2, oper)
					if err != nil {
						return -1, err
					}
					numStack.Push(res)
					operStack.Push(tempVal)
				} else {
					operStack.Push(tempVal)
				}
			}
		} else {
			i := idx
			for i < length && !isOper(int(str[i])) {
				i++
			}
			tempVal, err := strconv.ParseInt(str[idx:i], 10, 64)
			if err != nil {
				return -1, err
			}
			numStack.Push(int(tempVal))
			idx = i - 1
		}
	}

	for !operStack.IsEmpty() {
		num2, err1 := numStack.Pop()
		num1, err2 := numStack.Pop()
		oper, err3 := operStack.Pop()
		if err1 != nil || err2 != nil || err3 != nil {
			return -1, errors.Join(err1, err2, err3)
		}
		res, err := cal(num1, num2, oper)
		if err != nil {
			return -1, err
		}
		numStack.Push(res)
	}
	return numStack.Pop()
}

func isOper(val int) bool {
	// +:43 -:45 *:42 /:47
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

func cal(num1 int, num2 int, oper int) (res int, err error) {
	switch oper {
	case 42:
		res = num1 * num2
	case 43:
		res = num1 + num2
	case 45:
		res = num1 - num2
	case 47:
		res = num1 / num2
	default:
		err = errors.New("unknown oper, only suppoer +-*/")
	}
	return
}

func priority(oper int) (pri int, err error) {
	if oper == 42 || oper == 47 {
		pri = 1
	} else if oper == 43 || oper == 45 {
		pri = 0
	} else {
		err = errors.New("unknown oper, only suppoer +-*/")
	}
	return
}
