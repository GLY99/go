package main

import "sort"

func maxProduct(n int) int {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return nums[0] * nums[1]
}
