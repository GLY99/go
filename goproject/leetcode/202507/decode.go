package main

// https://leetcode.cn/problems/decode-xored-array/?envType=problem-list-v2&envId=array

func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	for i, e := range encoded {
		ans[i+1] = ans[i] ^ e
	}
	return ans
}
