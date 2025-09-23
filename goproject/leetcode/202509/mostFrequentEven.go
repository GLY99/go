package main

import "math"

// https://leetcode.cn/problems/most-frequent-even-element/?envType=problem-list-v2&envId=array

func mostFrequentEven(nums []int) int {
	counter := make(map[int]int)
	maxCount := 0
	ans := math.MaxInt64
	for _, num := range nums {
		if num%2 == 0 {
			counter[num]++
			if counter[num] > maxCount {
				maxCount = counter[num]
				ans = num
			} else if counter[num] == maxCount {
				if num < ans {
					ans = num
				}
			}
		}
	}
	if len(counter) == 0 {
		return -1
	}
	return ans
}
