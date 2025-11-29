package main

// https://leetcode.cn/problems/minimum-operations-to-make-array-sum-divisible-by-k/?envType=daily-question&envId=2025-11-29

func minOperations(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum % k
}
