package main

// https://leetcode.cn/problems/count-good-numbers/?envType=daily-question&envId=2025-04-13

func countGoodNumbers(n int64) int {
	mod := int64(1000000007)
	quickmul := func(x, y int64) int64 {
		var ret int64 = 1
		mul := x    // 1
		for y > 0 { // y - 1
			if y%2 == 1 {
				ret = ret * mul % mod
			}
			mul = mul * mul % mod
			y /= 2
		}
		return ret
	}
	return int(quickmul(5, (n+1)/2) * quickmul(4, n/2))
}
