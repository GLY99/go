package main

// https://leetcode.cn/problems/count-vowel-substrings-of-a-string/?envType=problem-list-v2&envId=hash-table

func countVowelSubstrings(word string) int {
	set := map[byte]struct{}{'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{}}
	n := len(word)
	res := 0
	for i := 0; i < n; i++ {
		newSet := map[byte]struct{}{}
		for j := i; j < n; j++ {
			newSet[word[j]] = struct{}{}
			if setEqual(set, newSet) {
				res++
			}
		}
	}
	return res
}

func setEqual[T comparable](a, b map[T]struct{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, _ := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}
