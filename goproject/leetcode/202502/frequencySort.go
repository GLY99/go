package main

import "sort"

// https://leetcode.cn/problems/sort-array-by-increasing-frequency/submissions/604097192/?envType=problem-list-v2&envId=hash-table

func frequencySort(nums []int) []int {
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num]++
	}
	var sFunc func(int, int) bool
	sFunc = func(i, j int) bool {
		return mapping[nums[i]] < mapping[nums[j]] || mapping[nums[i]] == mapping[nums[j]] && nums[i] > nums[j]
	}
	sort.Slice(nums, sFunc)
	return nums
}
