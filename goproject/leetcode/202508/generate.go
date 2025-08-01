package main

// https://leetcode.cn/problems/pascals-triangle/description/?envType=daily-question&envId=2025-08-01

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	ans := make([][]int, 0)
	ans = append(ans, []int{1})
	ans = append(ans, []int{1, 1})
	for i := 2; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		tmp[i] = 1
		for j := 1; j < i; j++ {
			tmp[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, tmp)
	}
	return ans
}
