package main

// https://leetcode.cn/problems/valid-sudoku/?envType=daily-question&envId=2025-08-30

func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int // 每行每列的hash
	var subBoxs [3][3][9]int // 每个小的9宫格hash
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			idx := c - '1'
			rows[i][idx]++
			cols[j][idx]++
			subBoxs[i/3][j/3][idx]++
			if rows[i][idx] > 1 || cols[j][idx] > 1 || subBoxs[i/3][j/3][idx] > 1 {
				return false
			}
		}
	}
	return true
}
