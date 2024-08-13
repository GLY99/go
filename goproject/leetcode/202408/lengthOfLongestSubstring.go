package main

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	left := 0
	right := 0
	mapping := map[byte]int{}
	maxSubLength := 0
	for right < length {
		// 一直移动右指针直到出现重复的字符
		for right < length && mapping[s[right]] == 0 {
			mapping[s[right]]++
			right++
		}
		// 出现一个不重复的长字符串
		maxSubLength = max(maxSubLength, right-left)
		if right >= length {
			break
		}
		// 移动左指针直到越过重复的字符
		for mapping[s[right]] != 0 {
			mapping[s[left]]--
			left++
		}
		// 此时左指针在重复字符的下一个位置，新的字符没有计数这里需要+1
		mapping[s[right]]++
		right++
	}
	return maxSubLength
}
