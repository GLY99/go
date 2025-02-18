package main

// https://leetcode.cn/problems/find-lucky-integer-in-an-array/?envType=problem-list-v2&envId=hash-table

func findLucky(arr []int) int {
	mapping := map[int]int{}
	for _, val := range arr {
		mapping[val]++
	}
	ans := -1
	for k, v := range mapping {
		if k == v {
			ans = max(ans, k)
		}
	}
	return ans
}
