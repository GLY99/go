package main

import "sort"

// https://leetcode.cn/problems/sort-vowels-in-a-string/?envType=daily-question&envId=2025-09-11

var vowels = map[rune]interface{}{
	'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{},
	'A': struct{}{}, 'E': struct{}{}, 'I': struct{}{}, 'O': struct{}{}, 'U': struct{}{},
}

func sortVowels(s string) string {
	var chars []rune
	for _, ch := range s {
		if _, ok := vowels[ch]; ok {
			chars = append(chars, ch)
		}
	}
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
	var result []rune
	idx := 0
	for _, ch := range s {
		if _, ok := vowels[ch]; ok {
			result = append(result, chars[idx])
			idx++
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}
