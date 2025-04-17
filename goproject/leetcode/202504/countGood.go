package main

// https://leetcode.cn/problems/count-the-number-of-good-subarrays/?envType=daily-question&envId=2025-04-16

func countGood(nums []int, k int) int64 {
	var ans int64 = 0
	count := make(map[int]int)
	same := 0
	n := len(nums)
	r := 0

	for l := 0; l < n; l++ {
		// 移动右指针，直到 same >= k 或 r 越界
		for same < k && r < n {
			same += count[nums[r]]
			count[nums[r]]++
			r++
		}
		// 如果 same >= k，则累加符合条件的子数组数目
		if same >= k {
			ans += int64(n - r + 1)
		}
		// 移动左指针，更新 same 和 count
		count[nums[l]]--
		same -= count[nums[l]]
	}
	return ans
}
