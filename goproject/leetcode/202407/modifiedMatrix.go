package main

func modifiedMatrix(matrix [][]int) [][]int {
	maxValMapping := make(map[int]int)
	for i := 0; i < len(matrix[0]); i++ {
		tmpMax := matrix[0][i]
		for j := 0; j < len(matrix); j++ {
			tmpMax = max(tmpMax, matrix[j][i])
		}
		maxValMapping[i] = tmpMax
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = maxValMapping[j]
			}
		}
	}
	return matrix
}
