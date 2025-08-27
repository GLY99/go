package main

// https://leetcode.cn/problems/container-with-most-water/?envType=problem-list-v2&envId=array

func maxArea(height []int) int {
	length := len(height)
	ans := 0
	left, right := 0, length-1
	for left < right {
		curArea := (right - left) * min(height[left], height[right])
		ans = max(ans, curArea)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
