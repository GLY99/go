package main

import "sort"

// https://leetcode.cn/problems/3sum/?envType=problem-list-v2&envId=array

func threeSum(nums []int) [][]int {
	// 从小到大排序，找出两个正数等于某个负数
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	n := len(nums)
	ans := make([][]int, 0)
	for a := 0; a < n; a++ {
		// 如果nums[a] > 0，那么必然不会存在 nums[a] + nums[b] + nums[c] == 0
		// 即-nums[a] == nums[b] + nums[c]
		if nums[a] > 0 {
			break
		}
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b := a + 1
		c := n - 1
		target := -1 * nums[a]
		for b < c {
			sum := nums[b] + nums[c]
			if sum == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
				b++
				c--
				// 去重；nums[a]每次是固定的，如果nums[b]和上次一样那么nums[c]也一定会一样，就会造成重复
				for b < c && nums[b] == nums[b-1] {
					b++
				}
				for b < c && nums[c] == nums[c+1] {
					c--
				}
			} else if sum > target {
				c--
			} else {
				b++
			}
		}
	}
	return ans
}
