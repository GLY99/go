package main

// https://leetcode.cn/problems/check-distances-between-same-letters/?envType=problem-list-v2&envId=array

func checkDistances(s string, distance []int) bool {
	idxs := make(map[rune]int)
	for idx, c := range s {
		if i, ok := idxs[c]; ok {
			if idx-i-1 != distance[c-'a'] {
				return false
			}
		} else {
			idxs[c] = idx
		}
	}
	return true
}
