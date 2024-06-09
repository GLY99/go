package main

func isAlienSorted(words []string, order string) bool {
	index := make([]int, 26)
	for idx, v := range order {
		index[v-'a'] = idx
	}
lable1:
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]) && j < len(words[i+1]); j++ {
			if index[words[i][j]-'a'] > index[words[i+1][j]-'a'] {
				return false
			} else if index[words[i][j]-'a'] < index[words[i+1][j]-'a'] {
				continue lable1
			}
		}
		if len(words[i]) > len(words[i+1]) {
			return false
		}
	}
	return true
}
