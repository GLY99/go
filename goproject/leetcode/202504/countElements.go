package main

import "math"

// https://leetcode.cn/problems/count-elements-with-strictly-smaller-and-greater-elements/?envType=problem-list-v2&envId=sorting

func countElements(nums []int) int {
	minNum := math.MaxInt64
	maxNum := math.MinInt64
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}
	ans := 0
	for _, num := range nums {
		if num > minNum && num < maxNum {
			ans++
		}
	}
	return ans
}
