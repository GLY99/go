package main

// https://leetcode.cn/problems/remove-one-element-to-make-the-array-strictly-increasing/?envType=problem-list-v2&envId=array

func canBeIncreasing(nums []int) bool {
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			return checkIncreasing(nums, i) || checkIncreasing(nums, i-1)
		}
	}
	return true
}

func checkIncreasing(nums []int, idx int) bool {
	n := len(nums)
	for i := 1; i < n-1; i++ {
		cur, pre := i, i-1
		if cur >= idx {
			cur += 1
		}
		if pre >= idx {
			pre += 1
		}
		if nums[cur] <= nums[pre] {
			return false
		}
	}
	return true
}
