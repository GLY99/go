package main

// https://leetcode.cn/problems/check-permutation-lcci/?envType=problem-list-v2&envId=sorting

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	map1, map2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		map1[s1[i]]++
		map2[s2[i]]++
	}
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if v2, ok := map2[k]; !ok {
			return false
		} else {
			if v2 != v {
				return false
			}
		}
	}
	return true
}
