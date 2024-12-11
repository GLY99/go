package main

func semiOrderedPermutation(nums []int) int {
	maxNum := nums[0]
	minNum := nums[0]
	maxIdx := 0
	minIdx := 0
	for idx, num := range nums {
		if num > maxNum {
			maxNum = num
			maxIdx = idx
		} else if num < minNum {
			minNum = num
			minIdx = idx
		}
	}
	if minIdx < maxIdx {
		return minIdx - 0 + len(nums) - maxIdx - 1
	} else {
		return minIdx - 0 + len(nums) - maxIdx - 2
	}
}
