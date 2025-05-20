package main

// https://leetcode.cn/problems/number-of-unequal-triplets-in-array/?envType=problem-list-v2&envId=sorting

func unequalTriplets(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	res, preCount, length := 0, 0, len(nums)
	for _, v := range count {
		res += v * preCount * (length - v - preCount)
		preCount += v
	}
	return res
}
