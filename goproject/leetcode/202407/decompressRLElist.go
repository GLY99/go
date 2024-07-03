package main

func decompressRLElist(nums []int) []int {
	length := len(nums)
	ans := []int{}
	for i := 0; i < length; i += 2 {
		for j := 0; j < nums[i]; j++ {
			ans = append(ans, nums[i+1])
		}
	}
	return ans
}
