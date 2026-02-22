package main

// https://leetcode.cn/problems/binary-gap/?envType=daily-question&envId=2026-02-22
func binaryGap(n int) int {
	preIdx := -1
	ans := -1
	idx := -1
	for n > 0 {
		idx++
		if n&1 != 0 && preIdx == -1 {
			preIdx = idx
		} else if n&1 != 0 {
			ans = max(idx-preIdx, ans)
			preIdx = idx
		}
		n = n >> 1
	}
	if ans == -1 {
		return 0
	}
	return ans
}
