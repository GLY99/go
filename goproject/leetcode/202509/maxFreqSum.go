package main

// https://leetcode.cn/problems/find-most-frequent-vowel-and-consonant/?envType=daily-question&envId=2025-09-13

func maxFreqSum(s string) int {
	counter := make([]int, 26)
	countA := 0
	countB := 0
	var vowels = map[rune]interface{}{
		'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{},
	}
	for _, char := range s {
		counter[char-'a']++
		if _, ok := vowels[char]; ok {
			countA = max(countA, counter[char-'a'])
		} else {
			countB = max(countB, counter[char-'a'])
		}
	}
	return countA + countB
}
