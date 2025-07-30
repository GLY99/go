package main

// https://leetcode.cn/problems/rotate-image/

// tmp = mat[row][col]
// mat[row][col] = mat[n - col - 1][row]
// mat[n - col - 1][row] = mat[n - row - 1][n - col - 1]
// mat[n - row - 1][n - col - 1] = mat[col][n - row - 1]
// mat[col][n - row - 1] = mat[row][col]

func rotate(mat [][]int) {
	n := len(mat)
	for row := 0; row < n/2; row++ {
		for col := 0; col < (n+1)/2; col++ {
			mat[row][col], mat[n-col-1][row], mat[n-row-1][n-col-1], mat[col][n-row-1] = mat[n-col-1][row], mat[n-row-1][n-col-1], mat[col][n-row-1], mat[row][col]
		}
	}
}
