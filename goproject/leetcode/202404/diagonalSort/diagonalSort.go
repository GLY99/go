package main

import "sort"

// https://leetcode.cn/problems/sort-the-matrix-diagonally/?envType=daily-question&envId=2024-04-29
func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	diag := make([][]int, m+n)

	// 将同一对角线的元素放入同一个数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diag[i-j+n] = append(diag[i-j+n], mat[i][j])
		}
	}

	// 对同一对角线的元素进行排序
	for _, arr := range diag {
		sort.Ints(arr)
	}

	// 排序后重新放回原数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = diag[i-j+n][0]
			diag[i-j+n] = diag[i-j+n][1:]
		}
	}
	return mat
}

func main() {
	diagonalSort([][]int{{1, 2}, {2, 1}})
}
