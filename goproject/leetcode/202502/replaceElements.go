package main

// https://leetcode.cn/problems/replace-elements-with-greatest-element-on-right-side/description/?envType=daily-question&envId=2025-02-16

func replaceElements(arr []int) []int {
	curMax := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = curMax
		curMax = max(curMax, tmp)
	}
	return arr
}
