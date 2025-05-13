package main

// https://leetcode.cn/problems/check-if-array-is-good/description/?envType=problem-list-v2&envId=sorting

func isGood(nums []int) bool {
	n := len(nums) - 1
	cnt := make([]int, len(nums))
	for _, num := range nums {
		if num > n || num == n && cnt[num] > 1 || num < n && cnt[num] > 0 {
			return false
		}
		cnt[num]++
	}
	return true
}
