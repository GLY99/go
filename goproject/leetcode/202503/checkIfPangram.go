package main

// https://leetcode.cn/problems/check-if-the-sentence-is-pangram/?envType=problem-list-v2&envId=hash-table

func checkIfPangram(sentence string) bool {
	x := 0
	for i := 0; i < len(sentence); i++ {
		x |= 1 << (sentence[i] - 'a')
	}
	return x == 1<<26-1
}
