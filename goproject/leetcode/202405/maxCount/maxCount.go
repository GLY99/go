package main

import "fmt"

// https://leetcode.cn/problems/range-addition-ii/
func maxCount(m int, n int, ops [][]int) int {
	mina := m
	minb := n
	for _, op := range ops {
		mina = min(mina, op[0])
		minb = min(minb, op[1])
	}
	return mina * minb
}

func main() {
	fmt.Println(maxCount(3, 3, [][]int{{1, 1}, {2, 2}}))
}
