package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	mapping := make(map[int]interface{})
	i := 0
	j := len(nums) - 1
	for i < j {
		mapping[nums[i]+nums[j]] = struct{}{}
		i++
		j--
	}
	return len(mapping)
}
