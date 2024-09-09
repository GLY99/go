package main

func isPrefixString(s string, words []string) bool {
	tempString := ""
	for i := 0; i < len(words); i++ {
		tempString += words[i]
		if len(tempString) > len(s) {
			return false
		}
		if tempString == s {
			return true
		}
	}
	return false
}
