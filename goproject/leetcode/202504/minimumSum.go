package main

import "sort"

// https://leetcode.cn/problems/minimum-sum-of-four-digit-number-after-splitting-digits/?envType=problem-list-v2&envId=sorting
func minimumSum(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return 10*(nums[0]+nums[1]) + (nums[2] + nums[3])
}
