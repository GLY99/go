package main

import "math"

// https://leetcode.cn/problems/row-with-maximum-ones/?envType=problem-list-v2&envId=array

func rowAndMaximumOnes(mat [][]int) []int {
	count := math.MinInt
	ans := 0
	for i, nums := range mat {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		if sum > count {
			count = sum
			ans = i
		}
	}
	return []int{ans, count}
}
