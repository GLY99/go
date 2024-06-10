package main

func numRookCaptures(board [][]byte) int {
	count, st, ed := 0, 0, 0
	for i, arr := range board {
		for j, v := range arr {
			if v == 'R' {
				st, ed = i, j
				break
			}
		}
	}
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		step := 0
		for {
			tx := st + step*dx[i]
			ty := ed + step*dy[i]
			if tx < 0 || ty < 0 || tx >= 8 || ty >= 8 || board[tx][ty] == 'B' {
				break
			}
			if board[tx][ty] == 'p' {
				count++
				break
			}
			step++
		}
	}
	return count
}
