package main

// https://leetcode.cn/problems/powx-n/description/

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	} else {
		return 1.0 / quickMul(x, -n)
	}
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := quickMul(x, n/2)
	if n%2 == 0 {
		return res * res
	}
	return res * res * x
}

func quickMul1(x float64, n int) float64 {
	res := 1.0
	mul := x
	// 当n的二进制最低位为1时才需要将结果记录到答案
	for n > 0 {
		if n%2 != 0 {
			res *= mul
		}
		mul *= mul
		n /= 2
	}
	return res
}
