package main

// https://leetcode.cn/problems/truncate-sentence/?envType=problem-list-v2&envId=array

func truncateSentence(s string, k int) string {
	sSlice := []byte(s)
	count := 0
	idx := 0
	for ; idx < len(sSlice); idx++ {
		if sSlice[idx] == ' ' {
			count++
		}
		if count == k {
			break
		}
	}
	if idx == len(sSlice) {
		return s
	}
	return string(sSlice[:idx])
}
