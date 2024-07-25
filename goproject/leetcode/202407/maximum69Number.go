package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	s := fmt.Sprintf("%d", num)
	sSlice := []byte(s)
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] == '6' {
			sSlice[i] = '9'
			break
		}
	}
	s = string(sSlice)
	ans, _ := strconv.Atoi(s)
	return ans
}
