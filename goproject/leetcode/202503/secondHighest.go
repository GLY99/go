package main

import "strconv"

// https://leetcode.cn/problems/second-largest-digit-in-a-string/?envType=problem-list-v2&envId=hash-table

func secondHighest(s string) int {
	max1 := -1
	max2 := -1
	for i := 0; i < len(s); i++ {
		num, err := strconv.Atoi(string(s[i]))
		if err == nil {
			if num > max1 {
				max2 = max1
				max1 = num
			} else if num < max1 && num > max2 {
				max2 = num
			}
		}
	}
	return max2
}
