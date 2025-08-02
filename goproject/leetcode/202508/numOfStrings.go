package main

// https://leetcode.cn/problems/number-of-strings-that-appear-as-substrings-in-word/?envType=problem-list-v2&envId=array

func numOfStrings(patterns []string, word string) int {
	ans := 0
	for _, pattern := range patterns {
		if check(pattern, word) {
			ans++
		}
	}
	return ans
}

func check(patterns string, word string) bool {
	n1 := len(patterns)
	n2 := len(word)
	for i := 0; i < n2-n1+1; i++ {
		flag := true
		for j := 0; j < n1; j++ {
			if patterns[j] != word[i+j] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}
