package main

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	prefixSum := 0
	for i := 0; i < len(nums); i++ {
		if sum-nums[i]-2*prefixSum == 0 {
			return i
		}
		prefixSum += nums[i]
	}
	return -1
}
