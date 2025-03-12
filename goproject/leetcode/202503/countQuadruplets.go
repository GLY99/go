package main

// https://leetcode.cn/problems/count-special-quadruplets/?envType=problem-list-v2&envId=hash-table

func countQuadruplets(nums []int) int {
	ans := 0
	mapping := map[int]int{}
	for b := len(nums) - 3; b >= 1; b-- {
		for _, d := range nums[b+2:] {
			mapping[d-nums[b+1]]++
		}
		for _, a := range nums[:b] {
			if sum := a + nums[b]; mapping[sum] > 0 {
				ans += mapping[sum]
			}
		}
	}
	return ans
}
