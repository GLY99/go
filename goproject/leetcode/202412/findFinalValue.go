package main

import "sort"

//https://leetcode.cn/problems/keep-multiplying-found-values-by-two/?envType=problem-list-v2&envId=simulation&difficulty=EASY&status=TO_DO

func findFinalValue(nums []int, original int) int {
	mapping := make(map[int]interface{}, 0)
	for _, num := range nums {
		mapping[num] = struct{}{}
	}
	for _, ok := mapping[original]; ok; _, ok = mapping[original] {
		original *= 2
	}
	return original
}

func findFinalValueI(nums []int, original int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for _, num := range nums {
		if original == num {
			original *= 2
		}
	}
	return original
}
