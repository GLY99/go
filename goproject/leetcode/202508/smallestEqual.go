package main

// https://leetcode.cn/problems/smallest-index-with-equal-value/?envType=problem-list-v2&envId=array

func smallestEqual(nums []int) int {
	for idx, num := range nums {
		if idx%10 == num {
			return idx
		}
	}
	return -1
}
