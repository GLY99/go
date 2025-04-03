package main

import "sort"

// https://leetcode.cn/problems/largest-number-after-digit-swaps-by-parity/?envType=problem-list-v2&envId=sorting

func largestInteger(num int) int {
	a, b, c := make([]int, 0), make([]int, 0), make([]int, 0)
	for ; num > 0; num /= 10 {
		if num%10%2 == 0 {
			a = append(a, num%10)
		} else {
			b = append(b, num%10)
		}
		c = append(c, num%10%2)
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	ans := 0
	j := 0
	k := 0
	for i := len(c) - 1; i >= 0; i-- {
		if c[i] == 0 {
			ans = ans*10 + a[j]
			j++
		} else {
			ans = ans*10 + b[k]
			k++
		}
	}
	return ans
}
