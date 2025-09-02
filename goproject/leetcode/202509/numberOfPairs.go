package main

// https://leetcode.cn/problems/maximum-number-of-pairs-in-array/?envType=problem-list-v2&envId=array

func numberOfPairs(nums []int) []int {
	mapping := [101]int{}
	for _, num := range nums {
		mapping[num]++
	}
	ans := make([]int, 2)
	for _, count := range mapping {
		ans[0] += count / 2
		ans[1] += count % 2
	}
	return ans
}
