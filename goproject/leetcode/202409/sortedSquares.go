package main

func sortedSquares(nums []int) []int {
	// 先找到第一个>=0的数字的位置
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	idx1 := left
	// 最后一个负数的位置
	idx2 := idx1 - 1

	var ans []int
	for idx2 >= 0 || idx1 < len(nums) {
		if idx2 >= 0 && idx1 < len(nums) {
			if -nums[idx2] <= nums[idx1] {
				ans = append(ans, nums[idx2]*nums[idx2])
				idx2--
			} else {
				ans = append(ans, nums[idx1]*nums[idx1])
				idx1++
			}
		} else if idx1 < len(nums) {
			ans = append(ans, nums[idx1]*nums[idx1])
			idx1++
		} else {
			ans = append(ans, nums[idx2]*nums[idx2])
			idx2--
		}
	}
	return ans
}
