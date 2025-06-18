package main

import "sort"

// https://leetcode.cn/problems/minimum-value-to-get-positive-step-by-step-sum/?envType=problem-list-v2&envId=array

func minStartValue(nums []int) int {
	sum, sumMinVal := 0, 0
	for _, num := range nums {
		sum += num
		sumMinVal = min(sumMinVal, sum)
	}
	return -sumMinVal + 1
}

func minStartValueI(nums []int) int {
	minVal := nums[0]
	for _, num := range nums {
		minVal = min(num, minVal)
	}
	return 1 + sort.Search(-minVal*len(nums), func(i int) bool {
		i++
		for _, num := range nums {
			i += num
			if i <= 0 {
				return false
			}
		}
		return true
	})
}
