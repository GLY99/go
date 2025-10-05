package main

import "strconv"

// https://leetcode.cn/problems/maximum-value-of-a-string-in-an-array/description/?envType=problem-list-v2&envId=array

func maximumValue(strs []string) int {
	res := 0
	for _, str := range strs {
		isDigits := true
		for _, c := range str {
			isDigits = isDigits && (c >= '0' && c <= '9')
		}
		if isDigits {
			num, _ := strconv.Atoi(str)
			res = max(res, num)
		} else {
			res = max(res, len(str))
		}
	}
	return res
}
