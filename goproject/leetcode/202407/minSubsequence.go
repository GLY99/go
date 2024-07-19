package main

import "sort"

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func minSubsequence(nums []int) []int {
	total := sum(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sum := 0
	for i, num := range nums {
		sum += num
		if sum > total-sum {
			return nums[:i+1]
		}
	}
	return nums
}
