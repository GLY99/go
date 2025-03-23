package main

import "sort"

func maxProductDifference(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return nums[0]*nums[1] - nums[len(nums)-1]*nums[len(nums)-2]
}

func maxProductDifference1(nums []int) int {
	length := len(nums)
	max1, max2 := max(nums[0], nums[1]), min(nums[0], nums[1])
	min1, min2 := min(nums[0], nums[1]), max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		num := nums[i]
		if num > max1 {
			max1, max2 = num, max1
		} else if num > max2 {
			max2 = num
		}
		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}
	}
	return max1*max2 - min1*min2
}
