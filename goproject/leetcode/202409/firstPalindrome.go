package main

func firstPalindrome(words []string) string {
lable1:
	for i := 0; i < len(words); i++ {
		l := 0
		r := len(words[i]) - 1
		for l < r {
			if words[i][l] != words[i][r] {
				continue lable1
			}
			l++
			r--
		}
		return words[i]
	}
	return ""
}
