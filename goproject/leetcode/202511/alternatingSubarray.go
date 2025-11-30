package main

// https://leetcode.cn/problems/longest-alternating-subarray/?envType=problem-list-v2&envId=array

func alternatingSubarray(nums []int) int {
	arr := make([]int, 0)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		arr = append(arr, nums[i+1]-nums[i])
	}
	pre := arr[0]
	counter := 0
	if pre == 1 {
		counter++
	}
	res := counter
	// 1 -1 1 -1 1 -1
	for i := 1; i < len(arr); i++ {
		if arr[i] == 1 {
			if pre == -1 {
				counter++
			} else {
				counter = 1
			}
		} else if arr[i] == -1 {
			if pre == 1 {
				counter++
			} else {
				counter = 0
			}
		}
		res = max(res, counter)
		pre = arr[i]
	}
	if res == 0 {
		return -1
	}
	return res + 1
}
