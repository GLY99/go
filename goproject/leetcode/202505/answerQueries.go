package main

import "sort"

// https://leetcode.cn/problems/longest-subsequence-with-limited-sum/?envType=problem-list-v2&envId=sorting

func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	n := len(nums)
	freq := make([]int, n)
	freq[0] = nums[0]
	for i := 1; i < n; i++ {
		freq[i] = freq[i-1] + nums[i]
	}
	ans := []int{}
	for _, q := range queries {
		res := binarySearch(freq, q)
		ans = append(ans, res+1)
	}
	return ans
}

func binarySearch(f []int, target int) int {
	left := 0
	right := len(f) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if f[mid] > target {
			right = mid - 1
		} else if f[mid] < target {
			left = mid + 1
		} else {
			right = mid
			break
		}
	}
	return right
}
