package main

// https://leetcode.cn/problems/most-frequent-number-following-key-in-an-array/?envType=problem-list-v2&envId=array

func mostFrequent(nums []int, key int) int {
	mapping := make(map[int]int)
	for idx := 0; idx < len(nums)-1; idx++ {
		if nums[idx] == key {
			mapping[nums[idx+1]]++
		}
	}
	ans := -1
	maxCount := -1
	for num, count := range mapping {
		if count > maxCount {
			ans = num
			maxCount = count
		}
	}
	return ans
}
