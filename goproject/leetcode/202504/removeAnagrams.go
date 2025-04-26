package main

// https://leetcode.cn/problems/find-resultant-array-after-removing-anagrams/?envType=problem-list-v2&envId=sorting

func removeAnagrams(words []string) []string {
	compare := func(idx int) bool {
		freq := make([]byte, 26)
		for _, word := range words[idx] {
			freq[byte(word)-'a']++
		}
		for _, word := range words[idx-1] {
			freq[byte(word)-'a']--
		}
		for _, val := range freq {
			if val != 0 {
				return false
			}
		}
		return true
	}
	ans := make([]string, 0)
	ans = append(ans, words[0])
	for i := 1; i < len(words); i++ {
		if compare(i) {
			continue
		}
		ans = append(ans, words[i])
	}
	return ans
}
