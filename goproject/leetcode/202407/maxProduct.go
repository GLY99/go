package main

import "sort"

func maxProduct1(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	return (nums[length-1] - 1) * (nums[length-2] - 1)
}

func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if a < b {
		a, b = b, a
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > a {
			a, b = nums[i], a
		} else if nums[i] > b {
			b = nums[i]
		}
	}
	return (a - 1) * (b - 1)
}
