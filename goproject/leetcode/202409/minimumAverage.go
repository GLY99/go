package main

import (
	"math"
	"sort"
)

func minimumAverage(nums []int) float64 {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	l := 0
	r := len(nums) - 1
	var res float64 = math.MaxFloat64
	var avg float64
	for l < r {
		avg = (float64(nums[l]) + float64(nums[r])) / 2.0
		res = math.Min(float64(avg), res)
		l++
		r--
	}
	return res
}
