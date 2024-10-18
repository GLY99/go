package main

import "sort"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findMaxK(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	left := 0
	right := len(nums) - 1
	for left < right {
		for left < right && nums[left] < 0 && nums[right] > 0 && abs(nums[left]) > nums[right] {
			left++
		}
		if nums[left] >= 0 || nums[right] <= 0 {
			break
		}
		if abs(nums[left]) == nums[right] {
			return nums[right]
		}
		for left < right && nums[left] < 0 && nums[right] > 0 && abs(nums[left]) < nums[right] {
			right--
		}
		if nums[left] >= 0 || nums[right] <= 0 {
			break
		}
		if abs(nums[left]) == nums[right] {
			return nums[right]
		}
	}
	return -1
}
