package main

// https://leetcode.cn/problems/kth-distinct-string-in-an-array/?envType=problem-list-v2&envId=array

func kthDistinct(arr []string, k int) string {
	mapping := make(map[string]bool)
	for _, str := range arr {
		_, ok := mapping[str]
		if ok {
			mapping[str] = false
		} else {
			mapping[str] = true
		}
	}
	for _, str := range arr {
		if mapping[str] {
			k--
			if k == 0 {
				return str
			}
		}
	}
	return ""
}
