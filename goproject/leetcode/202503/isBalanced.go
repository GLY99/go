package main

import "strconv"

// https://leetcode.cn/problems/check-balanced-string/?envType=daily-question&envId=2025-03-14

func isBalanced(num string) bool {
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(num); i += 2 {
		n1, _ := strconv.Atoi(string(num[i]))
		sum1 += n1
		if i+1 < len(num) {
			n2, _ := strconv.Atoi(string(num[i+1]))
			sum2 += n2
		}
	}
	return sum1 == sum2
}
