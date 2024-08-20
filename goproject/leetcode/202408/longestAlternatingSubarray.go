package main

func longestAlternatingSubarray(nums []int, threshold int) int {
	ans := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i]%2 != 0 {
			continue
		}
		if nums[i] > threshold {
			continue
		}
		j := i + 1
		for j < length && nums[j]%2 != nums[j-1]%2 && nums[j] <= threshold {
			j++
		}
		ans = max(j-i, ans)
		i = j - 1
	}
	return ans
}
