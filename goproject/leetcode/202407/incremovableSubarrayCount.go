package main

func incremovableSubarrayCount(nums []int) int {
	length := len(nums)
	ans := 0
	for left := 0; left < length; left++ {
		for right := left; right < length; right++ {
			if isIncrease(nums, left, right) {
				ans++
			}
		}
	}
	return ans
}

func isIncrease(nums []int, left int, right int) (ret bool) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if i >= left-1 && i <= right {
			continue
		}
		if nums[i] >= nums[i+1] {
			ret = false
			return
		}
	}
	if left-1 >= 0 && right < length-1 && nums[left-1] >= nums[right+1] {
		ret = false
		return
	}
	ret = true
	return
}
