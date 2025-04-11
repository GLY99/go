package main

// https://leetcode.cn/problems/count-symmetric-integers/?envType=daily-question&envId=2025-04-11

func countSymmetricIntegers(low int, high int) int {
	res := 0
	for i := low; i <= high; i++ {
		if i < 100 && i%11 == 0 {
			res++
		} else if i > 1000 && i < 10000 {
			if i/1000+(i/100)%10 == i%10+(i/10)%10 {
				res++
			}
		}
	}
	return res
}
