package main

// https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/submissions/600318047/?envType=daily-question&envId=2025-02-17

func findSpecialInteger(arr []int) int {
	cur, count := arr[0], 0
	for _, val := range arr {
		if cur == val {
			count++
			if count*4 > len(arr) {
				return cur
			}
		} else {
			cur, count = val, 1
		}
	}
	return -1
}
