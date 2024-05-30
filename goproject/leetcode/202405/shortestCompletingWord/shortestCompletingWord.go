package main

import "fmt"

func shortestCompletingWord(licensePlate string, words []string) string {
	collectLetters := make([]int, 26)
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] >= 65 && licensePlate[i] <= 90 {
			collectLetters[licensePlate[i]-'A']++
		} else if licensePlate[i] >= 97 && licensePlate[i] <= 122 {
			collectLetters[licensePlate[i]-'a']++
		}
	}
	var res string
lable1:
	for _, word := range words {
		chars := make([]int, 26)
		for i := 0; i < len(word); i++ {
			if word[i] >= 65 && word[i] <= 90 {
				chars[word[i]-'A']++
			} else if word[i] >= 97 && word[i] <= 122 {
				chars[word[i]-'a']++
			}
		}
		for i := 0; i < 26; i++ {
			if chars[i] < collectLetters[i] {
				continue lable1
			}
		}
		if res == "" || len(word) < len(res) {
			res = word
		}
	}
	return res
}

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
}
