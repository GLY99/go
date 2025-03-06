package main

// https://leetcode.cn/problems/check-if-all-the-integers-in-a-range-are-covered/?envType=problem-list-v2&envId=hash-table

func isCovered(ranges [][]int, left int, right int) bool {
	mapping := make(map[int]interface{})
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			mapping[i] = struct{}{}
		}
	}
	for i := left; i <= right; i++ {
		if _, ok := mapping[i]; !ok {
			return ok
		}
	}
	return true
}
