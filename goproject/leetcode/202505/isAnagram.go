package main

// https://leetcode.cn/problems/dKk3P7/?envType=problem-list-v2&envId=sorting

func isAnagram(s string, t string) bool {
	map1 := make([]int, 26)
	map2 := make([]int, 26)
	l1, l2 := len(s), len(t)
	if l1 != l2 || s == t {
		return false
	}
	for i := 0; i < l1; i++ {
		map1[s[i]-'a']++
		map2[t[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if map1[i] != map2[i] {
			return false
		}
	}
	return true
}
