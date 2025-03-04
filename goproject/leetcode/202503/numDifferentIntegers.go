package main

// https://leetcode.cn/problems/number-of-different-integers-in-a-string/description/?envType=problem-list-v2&envId=hash-table

func numDifferentIntegers(word string) int {
	mapping := make(map[string]interface{})
	s, e := 0, 0
	for s < len(word) && e < len(word) {
		for s < len(word) && word[s] >= 'a' && word[s] <= 'z' {
			s++
		}
		if s >= len(word) {
			break
		}
		e = s
		for e < len(word) && word[e] >= '0' && word[e] <= '9' {
			e++
		}
		numStr := word[s:e]
		zeroIdx := 0
		for zeroIdx < len(numStr) {
			if numStr[zeroIdx] != '0' {
				break
			}
			zeroIdx++
		}
		numStr = numStr[zeroIdx:]
		mapping[numStr] = struct{}{}
		s = e
	}
	return len(mapping)
}
