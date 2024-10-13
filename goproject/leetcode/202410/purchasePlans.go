package main

import "sort"

func purchasePlans(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	j := len(nums) - 1
	ans := 0
	for i := 0; i < len(nums); i++ {
		for ; j > i; j-- {
			if nums[i]+nums[j] <= target {
				break
			}
		}
		if j > i {
			ans += j - i
		}
	}
	return ans % 1000000007
}
