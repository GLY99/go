package main

func maximumTripletValue(nums []int) int64 {
	var res int64 = 0
	for k := 2; k < len(nums); k++ {
		m := nums[0]
		for j := 1; j < k; j++ {
			res = max(res, int64(m-nums[j])*int64(nums[k]))
			m = max(m, nums[j])
		}
	}
	return res
}

func maximumTripletValue1(nums []int) int64 {
	length := len(nums)
	leftMax, rightMax := make([]int, length), make([]int, length)
	for i := 1; i < len(nums); i++ {
		leftMax[i] = max(leftMax[i-1], nums[i-1])
		rightMax[length-1-i] = max(rightMax[length-i], nums[length-i])
	}
	var ans int64 = 0
	for j := 1; j < len(nums)-1; j++ {
		ans = max(ans, int64(leftMax[j]-nums[j])*int64(rightMax[j]))
	}
	return ans
}
