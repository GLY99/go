package main

func countGoodSubstrings(s string) int {
	length := len(s)
	mapping := map[byte]int{}
	ans := 0
	for i := 0; i < length; i++ {
		// 每次加入最右边的字符
		mapping[s[i]]++
		// 在加入第四个字符时需要移除最左边的第一个字符
		if i >= 3 {
			if mapping[s[i-3]] > 1 {
				mapping[s[i-3]]--
			} else if mapping[s[i-3]] == 1 {
				delete(mapping, s[i-3])
			}
		}
		// 如果mapping长度是3则是一个答案
		if len(mapping) == 3 {
			ans++
		}
	}
	return ans
}
