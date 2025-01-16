package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	res := math.MaxInt
	for i, v := range nums {
		if v >= k {
			return 1
		}
		value := 0
		for j := i; j >= 0; j-- {
			value |= nums[j]
			if value >= k {
				res = min(res, i-j+1)
				break
			}
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
