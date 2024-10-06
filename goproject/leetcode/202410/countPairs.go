package main

import "sort"

func countPairs(nums []int, target int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				ans++
			}
		}
	}
	return ans
}

func countPairs1(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	ans := 0
	for i, j := 0, len(nums)-1; i < j; i++ {
		for i < j && nums[i]+nums[j] >= target {
			j--
		}
		ans += j - i
	}
	return ans
}
