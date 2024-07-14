package main

func getLongestSubsequence(words []string, groups []int) []string {
	ans := []string{}
	length := len(words)
	for i, x := range groups {
		if i == length-1 || x != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return ans
}
