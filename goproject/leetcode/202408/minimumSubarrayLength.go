package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	res := math.MaxInt
	for i, v := range nums {
		if v >= k {
			return 1
		}
		for j := i - 1; j >= 0; j-- {
			if v|nums[j] == nums[j] {
				break
			}
			nums[j] |= v
			if nums[j] >= k {
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
