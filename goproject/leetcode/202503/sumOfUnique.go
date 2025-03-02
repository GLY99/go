package main

// https://leetcode.cn/problems/sum-of-unique-elements/?envType=problem-list-v2&envId=hash-table

func sumOfUnique(nums []int) int {
	mapping := map[int]int{}
	for _, num := range nums {
		mapping[num]++
	}
	res := 0
	for k, v := range mapping {
		if v == 1 {
			res += k
		}
	}
	return res
}
