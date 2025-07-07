package main

// https://leetcode.cn/problems/richest-customer-wealth/?envType=problem-list-v2&envId=array

func maximumWealth(accounts [][]int) int {
	ans := 0
	for i := 0; i < len(accounts); i++ {
		tmp := 0
		for j := 0; j < len(accounts[i]); j++ {
			tmp += accounts[i][j]
		}
		ans = max(ans, tmp)
	}
	return ans
}
