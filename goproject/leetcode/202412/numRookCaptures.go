package main

// https://leetcode.cn/problems/available-captures-for-rook/?envType=daily-question&envId=2024-12-06

func numRookCaptures(board [][]byte) int {
	count, x0, y0 := 0, 0, 0
	for x, items := range board {
		for y, item := range items {
			if item == 'R' {
				x0, y0 = x, y
				break
			}
		}
	}
	steps := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 0; i < 4; i++ {
		x1, y1 := x0, y0
		for {
			x1, y1 = x1+steps[i][0], y1+steps[i][1]
			if x1 < 0 || x1 >= 8 || y1 < 0 || y1 >= 8 || board[x1][y1] == 'B' {
				break
			}
			if board[x1][y1] == 'p' {
				count++
				break
			}
		}
	}
	return count
}
