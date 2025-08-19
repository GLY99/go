package main

// https://leetcode.cn/problems/divide-array-into-equal-pairs/?envType=problem-list-v2&envId=array

func divideArray(nums []int) bool {
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num]++
	}
	for _, count := range mapping {
		if count%2 != 0 {
			return false
		}
	}
	return true
}
