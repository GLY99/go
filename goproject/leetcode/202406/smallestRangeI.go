package main

func smallestRangeI(nums []int, k int) int {
	// min + x = max - y
	// ans = max - min - x - y
	// ans = max - min - (x + y)
	// 0 <= x + y <= 2k
	// ans最小的时候是x + y = 2k
	// ans = max - min - 2k
	// y = -2x + b
	maxVal := nums[0]
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		} else if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	ans := maxVal - minVal - 2*k
	if ans > 0 {
		return ans
	}
	return 0
}
