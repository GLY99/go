package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Sort(sort.IntSlice(nums))
	ans := math.MaxInt64
	for i, num := range nums[:len(nums)-k+1] {
		ans = min(ans, nums[i+k-1]-num)
	}
	return ans
}
