package main

// https://leetcode.cn/problems/maximum-sum-with-exactly-k-elements/?envType=problem-list-v2&envId=array

func maximizeSum(nums []int, k int) int {
	maxVal := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return (2*maxVal + k - 1) * k / 2
}
