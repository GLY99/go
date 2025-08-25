package main

// https://leetcode.cn/problems/diagonal-traverse/?envType=daily-question&envId=2025-08-25

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0)
	x, y := 0, 0
	d := -1
	// 对角线从0到m+n-2
	for i := 0; i < m+n-1; i++ {
		for x >= 0 && x < m && y >= 0 && y < n {
			ans = append(ans, mat[x][y])
			// 每个对角线的x+y=i; 所以通过控制d的正负来确定走向，得出x，最后计算出y
			x += d
			y = i - x
		}
		// 越界后退回来方便计算新的起点
		x -= d
		y = i - x
		if i%2 == 0 {
			// 向右上走
			// 如果是走在非最长的对角线上那么右移后就是新的起点
			// 如果是走在最长的对角线上那么向下才是新起点
			if y == n-1 {
				x++
			} else {
				y++
			}
		} else {
			// 向左下走
			// 如果是走在非最长对角线向下是新起点
			// 如果走在最长对角线向右是新起点
			if x == m-1 {
				y++
			} else {
				x++
			}
		}
		// 换方向
		d *= -1
	}
	return ans
}
