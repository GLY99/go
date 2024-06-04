package main

func shortestToChar(s string, c byte) []int {
	length := len(s)
	ans := make([]int, length)
	idx := -length
	// 从右边遍历，寻找每个字符距离左边最近的字符c的距离
	for i := 0; i < length; i++ {
		if s[i] == c {
			idx = i
		}
		ans[i] = i - idx
	}
	// 从左边遍历，寻找每个字符距离右边最近的字符c的距离
	idx = length * 2
	for i := length - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}
