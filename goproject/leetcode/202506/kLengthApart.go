package main

// https://leetcode.cn/problems/check-if-all-1s-are-at-least-length-k-places-away/?envType=problem-list-v2&envId=array

func kLengthApart(nums []int, k int) bool {
	preIdx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 && preIdx == -1 {
			preIdx = i
			continue
		} else if nums[i] == 1 && preIdx != -1 {
			if i-preIdx-1 < k {
				return false
			} else {
				preIdx = i
			}
		}
	}
	return true
}
