package main

// https://leetcode.cn/problems/count-the-number-of-consistent-strings/?envType=problem-list-v2&envId=hash-table

func countConsistentStrings(allowed string, words []string) int {
	var mapping map[byte]interface{} = make(map[byte]interface{})
	for i := 0; i < len(allowed); i++ {
		mapping[allowed[i]] = struct{}{}
	}
	count := 0
lable1:
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if _, ok := mapping[word[i]]; !ok {
				continue lable1
			}
		}
		count++
	}
	return count
}

func countConsistentStrings1(allowed string, words []string) int {
	mask := 0
	for _, c := range allowed {
		mask |= 1 << (c - 'a')
	}
	count := 0
	for _, word := range words {
		curmask := 0
		for _, c := range word {
			curmask |= 1 << (c - 'a')
		}
		if curmask|mask == mask {
			count++
		}
	}
	return count
}
