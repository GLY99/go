package main

func constructTransformedArray(nums []int) []int {
	var result []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = nums[((i+nums[i])%len(nums)+len(nums))%len(nums)]
	}
	return result
}
