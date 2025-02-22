package main

// https://leetcode.cn/problems/count-pairs-of-similar-strings/submissions/602090433/?envType=daily-question&envId=2025-02-22

func similarPairs(words []string) int {
	res := 0
	cnt := make(map[int]int)
	for _, word := range words {
		state := 0
		for _, c := range word {
			state |= 1 << (c - 'a')
		}
		res += cnt[state]
		cnt[state]++
	}
	return res
}
