package main

// https://leetcode.cn/problems/matrix-diagonal-sum/?envType=problem-list-v2&envId=array

func diagonalSum(mat [][]int) int {
	n := len(mat)
	mid := n / 2
	ans := 0
	for i := 0; i < n; i++ {
		ans += mat[i][i] + mat[i][n-i-1]
	}
	if n%2 == 1 {
		ans -= mat[mid][mid]
	}
	return ans
}
