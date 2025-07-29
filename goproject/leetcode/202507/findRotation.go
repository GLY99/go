package main

// https://leetcode.cn/problems/determine-whether-matrix-can-be-obtained-by-rotation/?envType=problem-list-v2&envId=array
// matrix[row][col] 旋转90后的位置
// matrix[col][n−row−1]=matrix[row][col]

// tmp = matrix[col][n−row−1]
// matrix[n−row−1][n−col−1]=matrix[col][n−row−1]

// tmp = matrix[n−row−1][n−col−1]
// matrix[n−row−1][n−col−1]=matrix[col][n−row−1]
// matrix[col][n−row−1]=matrix[row][col]

// tmp = matrix[n−col−1][row]
// matrix[n−col−1][row] = matrix[n−row−1][n−col−1]
// matrix[n−row−1][n−col−1]=matrix[col][n−row−1]
// matrix[col][n−row−1]=matrix[row][col]

// tmp = matrix[row][col]
// matrix[row][col] = matrix[n−col−1][row]
// matrix[n−col−1][row] = matrix[n−row−1][n−col−1]
// matrix[n−row−1][n−col−1]=matrix[col][n−row−1]
// matrix[col][n−row−1]=matrix[row][col]

func findRotation(mat [][]int, target [][]int) bool {
	length := len(mat)
	for k := 0; k < 4; k++ {
		for i := 0; i < length/2; i++ {
			for j := 0; j < (length+1)/2; j++ {
				mat[i][j], mat[length-j-1][i], mat[length-i-1][length-j-1], mat[j][length-i-1] = mat[length-j-1][i], mat[length-i-1][length-j-1], mat[j][length-i-1], mat[i][j]
			}
		}
		if sliceEqual(mat, target) {
			return true
		}
	}
	return false
}

func sliceEqual(nums [][]int, target [][]int) bool {
	if len(nums) != len(target) {
		return false
	}
	for i := range nums {
		if len(nums[i]) != len(target[i]) {
			return false
		}
		for j := range nums[i] {
			if nums[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}
