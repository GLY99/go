package main

// https://leetcode.cn/problems/special-positions-in-a-binary-matrix/?envType=problem-list-v2&envId=array

func numSpecial(mat [][]int) int {
	rowsSum := make([]int, len(mat))
	colsSum := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			rowsSum[i] += mat[i][j]
			colsSum[j] += mat[i][j]
		}
	}
	ans := 0
	for i, row := range mat {
		for j, x := range row {
			if x == 1 && rowsSum[i] == 1 && colsSum[j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func numSpecial1(mat [][]int) int {
	for i, row := range mat {
		cnt := 0
		for _, x := range row {
			cnt += x
		}
		if i == 0 {
			cnt--
		}
		if cnt > 0 {
			for j, x := range row {
				if x == 1 {
					mat[0][j] += cnt
				}
			}
		}
	}
	ans := 0
	for _, x := range mat[0] {
		if x == 1 {
			ans++
		}
	}
	return ans
}
