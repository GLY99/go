package main

// https://leetcode.cn/problems/count-hills-and-valleys-in-an-array/?envType=daily-question&envId=2025-07-27

func countHillValley(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	count := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		l := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				l++
				break
			} else if nums[i] < nums[j] {
				l--
				break
			}
		}
		r := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				r++
				break
			} else if nums[i] < nums[j] {
				r--
				break
			}
		}
		if l != 0 && l == r {
			count++
		}
	}
	return count
}
