package main

// https://leetcode.cn/problems/number-of-even-and-odd-bits/submissions/601467375/?envType=daily-question&envId=2025-02-20

func evenOddBit(n int) []int {
	ans := []int{0, 0}
	idx := 0
	for n != 0 {
		if n&1 == 1 {
			ans[idx%2]++
		}
		idx++
		n >>= 1
	}
	return ans
}
