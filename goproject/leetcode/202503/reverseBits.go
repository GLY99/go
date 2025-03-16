package main

// https://leetcode.cn/problems/reverse-bits-lcci/?envType=problem-list-v2&envId=dynamic-programming

func reverseBits(num int) int {
	cur := 0
	insert := 0
	ans := 1
	for i := 0; i < 32; i++ {
		if (num & (1 << i)) != 0 {
			cur += 1
			insert += 1
		} else {
			insert = cur + 1
			cur = 0
		}
		ans = max(ans, insert)
	}
	return ans
}
