package main

func isPrefixOfWord(sentence string, searchWord string) int {
	speaceCount := 0
	for i := 0; i <= len(sentence)-len(searchWord); i++ {
		if string(sentence[i:i+len(searchWord)]) == searchWord {
			return speaceCount + 1
		} else {
			for i < len(sentence) && sentence[i] != ' ' {
				i++
			}
			speaceCount++
		}
	}
	return -1
}
