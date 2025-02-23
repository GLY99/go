package main

// https://leetcode.cn/problems/number-of-good-pairs/?envType=problem-list-v2&envId=hash-table

func numIdenticalPairs(nums []int) int {
	mapping := make(map[int]int)
	ans := 0
	for _, num := range nums {
		ans += mapping[num]
		mapping[num]++
	}
	return ans
}
