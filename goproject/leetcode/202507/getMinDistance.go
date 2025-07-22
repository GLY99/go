package main

import "math"

// https://leetcode.cn/problems/minimum-distance-to-the-target-element/?envType=problem-list-v2&envId=array

func getMinDistance(nums []int, target int, start int) int {
	ans := math.MaxInt
	for idx, num := range nums {
		if num == target {
			ans = min(ans, abs(idx-start))
		}
	}
	return ans
}
