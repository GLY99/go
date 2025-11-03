package main

// https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range/?envType=problem-list-v2&envId=array

func vowelStrings(words []string, left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		n := len(words[i])
		if isVowel(words[i][0]) && isVowel(words[i][n-1]) {
			ans++
		}
	}

	return ans

}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}
