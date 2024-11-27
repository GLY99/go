package main

func buildArray(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
