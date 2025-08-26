package main

// https://leetcode.cn/problems/maximum-area-of-longest-diagonal-rectangle/?envType=daily-question&envId=2025-08-26

func areaOfMaxDiagonal(dimensions [][]int) int {
	ans := dimensions[0][0] * dimensions[0][1]
	maxDiagonal := dimensions[0][0]*dimensions[0][0] + dimensions[0][1]*dimensions[0][1]
	for _, dimension := range dimensions {
		curDiagonal := dimension[0]*dimension[0] + dimension[1]*dimension[1]
		curArea := dimension[0] * dimension[1]
		if curDiagonal > maxDiagonal {
			maxDiagonal = curDiagonal
			ans = curArea
		} else if curDiagonal == maxDiagonal {
			if curArea > ans {
				ans = curArea
			}
		}
	}
	return ans
}
