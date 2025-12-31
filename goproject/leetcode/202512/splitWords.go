package main

// https://leetcode.cn/problems/split-strings-by-separator/?envType=problem-list-v2&envId=array

func splitWordsBySeparator(words []string, separator byte) []string {
	ans := make([]string, 0)
	for _, word := range words {
		start, end := 0, 0
		n := len(word)
		for start < n && end < n {
			for start < n && word[start] == separator {
				start++
			}
			if start >= n {
				break
			}
			end = start
			for end < n && word[end] != separator {
				end++
			}
			if end >= n {
				ans = append(ans, word[start:n])
			} else {
				ans = append(ans, word[start:end])
			}
			start = end + 1
		}
	}
	return ans
}
