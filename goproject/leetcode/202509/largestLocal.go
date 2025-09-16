package main

// https://leetcode.cn/problems/largest-local-values-in-a-matrix/?envType=problem-list-v2&envId=array

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		ans[i] = make([]int, n-2)
		for j := 0; j < n-2; j++ {
			curMax := ans[i][j]
			for a := i; a < i+3; a++ {
				for b := j; b < j+3; b++ {
					curMax = max(curMax, grid[a][b])
				}
			}
			ans[i][j] = curMax
		}
	}
	return ans
}
