package main

import (
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/split-with-minimum-sum/?envType=problem-list-v2&envId=sorting

func splitNum(num int) int {
	strNum := []byte(strconv.Itoa(num))
	sort.Slice(strNum, func(i, j int) bool { return strNum[i] < strNum[j] })
	num1, num2 := 0, 0
	for i := 0; i < len(strNum); i++ {
		if i%2 == 0 {
			num1 = num1*10 + int(strNum[i]-'0')
		} else {
			num2 = num2*10 + int(strNum[i]-'0')
		}
	}
	return num1 + num2
}
