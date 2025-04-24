package main

import "sort"

// https://leetcode.cn/problems/intersection-of-multiple-arrays/?envType=problem-list-v2&envId=sorting

func intersection(nums [][]int) []int {
	mapping := make(map[int]int)
	for _, tmpNums := range nums {
		for _, num := range tmpNums {
			mapping[num]++
		}
	}
	ans := make([]int, 0)
	for k, v := range mapping {
		if v == len(nums) {
			ans = append(ans, k)
		}
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })
	return ans
}
