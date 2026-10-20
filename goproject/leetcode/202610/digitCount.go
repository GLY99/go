package main

// https://leetcode.cn/problems/check-if-number-has-equal-digit-count-and-digit-value/?envType=problem-list-v2&envId=hash-table

func digitCount(num string) bool {
	n := len(num)
	mapping := make([]int, 10)
	for i := 0; i < n; i++ {
		mapping[int(num[i]-'0')]++
	}
	for i := 0; i < n; i++ {
		if mapping[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
