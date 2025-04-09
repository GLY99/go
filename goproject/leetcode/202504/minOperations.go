package main

// https://leetcode.cn/problems/minimum-operations-to-make-array-values-equal-to-k/?envType=daily-question&envId=2025-04-09

func minOperations(nums []int, k int) int {
	set := make(map[int]interface{})
	for _, num := range nums {
		if num < k {
			return -1
		} else if num > k {
			set[num] = struct{}{}
		}
	}
	return len(set)
}
