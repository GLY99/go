package main

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			flag1, flag2 := true, true
			// 删除left位置的字符，判断剩下的字符串是不是回文
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			// 删除rigth位置字符，判断剩下的字符串是不是回文
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
