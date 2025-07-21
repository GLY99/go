package main

// https://leetcode.cn/problems/delete-characters-to-make-fancy-string/?envType=daily-question&envId=2025-07-21

func makeFancyString(s string) string {
	res := []rune{}
	for _, ch := range s {
		n := len(res)
		if n >= 2 && res[n-1] == ch && res[n-2] == ch {
			continue
		}
		res = append(res, ch)
	}
	return string(res)
}
