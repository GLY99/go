package main

func maxOperations(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}
	count := 0
	sum := nums[0] + nums[1]
	for i := 0; i < length-1; i = i + 2 {
		if nums[i]+nums[i+1] == sum {
			count++
		} else {
			break
		}
	}
	return count
}
