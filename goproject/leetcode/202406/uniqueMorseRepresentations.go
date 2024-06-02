package main

func uniqueMorseRepresentations(words []string) int {
	var morse = []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
		"....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-",
		"...-", ".--", "-..-", "-.--", "--..",
	}
	set := map[string]struct{}{}
	for _, word := range words {
		trans := ""
		for _, char := range word {
			trans += morse[char-'a']
		}
		set[trans] = struct{}{}
	}
	return len(set)
}
