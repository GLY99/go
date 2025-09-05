package main

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/3sum-closest/?envType=problem-list-v2&envId=array

func threeSumClosest(nums []int, target int) int {
	best := math.MaxInt32
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	updater := func(num int) {
		if abs(num, target) < abs(best, target) {
			best = num
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			updater(sum)
			if sum > target {
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return best
}
