package main

import "math"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findClosestNumber(nums []int) int {
	res := math.MaxInt64
	for _, num := range nums {
		if num == 0 {
			return num
		}
		absNum := abs(num)
		absRes := abs(res)
		if absNum < absRes {
			res = num
		} else if absNum == absRes {
			if num > res {
				res = num
			}
		}
	}
	return res
}
