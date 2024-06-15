package main

func countCharacters(words []string, chars string) int {
	charsMapping := make(map[rune]int)
	for _, char := range chars {
		charsMapping[char]++
	}
	res := 0
lable1:
	for _, word := range words {
		wordCharsMapping := make(map[rune]int)
		for _, char := range word {
			wordCharsMapping[char]++
			if wordCharsMapping[char] > charsMapping[char] {
				continue lable1
			}
		}
		res += len(word)
	}
	return res
}
