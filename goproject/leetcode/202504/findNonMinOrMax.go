package main

import "math"

// https://leetcode.cn/problems/neither-minimum-nor-maximum/?envType=problem-list-v2&envId=sorting

func findNonMinOrMax(nums []int) int {
	minValue := math.MaxInt64
	maxValue := math.MinInt64
	for _, num := range nums {
		if num < minValue {
			minValue = num
		}
		if num > maxValue {
			maxValue = num
		}
	}
	for _, num := range nums {
		if num > minValue && num < maxValue {
			return num
		}
	}
	return -1
}
