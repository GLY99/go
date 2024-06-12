package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	ans := 0
	sort.Ints(nums)
	for idx, num := range nums {
		if num <= 0 && k > 0 {
			nums[idx] = -num
			k--
		}
		ans += nums[idx]
	}
	sort.Ints(nums)
	if k%2 == 0 {
		return ans
	} else {
		return ans - 2*nums[0]
	}
}
