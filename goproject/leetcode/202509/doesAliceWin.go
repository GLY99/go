package main

// https://leetcode.cn/problems/vowels-game-in-a-string/?envType=daily-question&envId=2025-09-12

func doesAliceWin(s string) bool {
	counter := 0
	var vowels = map[rune]interface{}{
		'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{},
	}
	for _, ch := range s {
		if _, ok := vowels[ch]; ok {
			counter++
		}
	}
	return counter != 0
}
