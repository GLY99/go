package main

import "sort"

func numberGame(nums []int) []int {
	ans := []int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums); i += 2 {
		ans = append(ans, nums[i+1])
		ans = append(ans, nums[i])
	}
	return ans
}
