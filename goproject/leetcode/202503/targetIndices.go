package main

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		if nums[i] == target {
			ans = append(ans, i)
		}
	}
	return ans
}
