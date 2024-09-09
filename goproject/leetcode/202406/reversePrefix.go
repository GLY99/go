package main

func reversePrefix(word string, ch byte) string {
	i := 0
	for ; i < len(word); i++ {
		if word[i] == ch {
			break
		}
	}
	if i == len(word) {
		return word
	}
	wordSlice := []byte(word)
	j := 0
	for j < i {
		wordSlice[i], wordSlice[j] = wordSlice[j], wordSlice[i]
		j++
		i--
	}
	return string(wordSlice)
}
