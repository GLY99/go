package main

func getLongestSubsequence(words []string, groups []int) []string {
	length := len(words)
	ans := make([]string, 0)
	for i, x := range groups {
		if i == length-1 || x != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return ans
}
