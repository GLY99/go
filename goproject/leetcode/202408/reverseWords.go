package main

func reverseWord(words []byte, start, end int) {
	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	start := 0
	end := 0
	length := len(s)
	words := []byte(s)
	for end < length {
		for end < length && words[end] != ' ' {
			end++
		}
		reverseWord(words, start, end-1)
		end++
		start = end
	}
	return string(words)
}
