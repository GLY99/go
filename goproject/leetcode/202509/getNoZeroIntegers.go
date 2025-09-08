package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/convert-integer-to-the-sum-of-two-no-zero-integers/?envType=daily-question&envId=2025-09-08

func getNoZeroIntegers(n int) []int {
	for i := 0; i <= n/2; i++ {
		j := n - i
		if !strings.Contains(strconv.Itoa(i), "0") && !strings.Contains(strconv.Itoa(j), "0") {
			return []int{i, j}
		}
	}
	return []int{}
}
