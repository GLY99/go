package main

import "sort"

// https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-i/?envType=daily-question&envId=2025-11-04

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	ans := make([]int, 0)
	for i := 0; i <= n-k; i++ {
		mapping := make(map[int]int, 0)
		s := make([]int, 0)
		for j := i; j < i+k; j++ {
			if _, ok := mapping[nums[j]]; !ok {
				s = append(s, nums[j])
			}
			mapping[nums[j]]++
		}
		sort.Slice(
			s,
			// i是位于后面的元素，j是位于前面的元素； 返回true时交换元素位置
			func(i, j int) bool {
				a := mapping[s[i]]
				b := mapping[s[j]]
				if a != b {
					return a > b
				} else {
					return s[i] > s[j]
				}
			},
		)
		sum := 0
		for j := 0; j < x && j < len(s); j++ {
			sum += s[j] * mapping[s[j]]
		}
		ans = append(ans, sum)
	}
	return ans
}
