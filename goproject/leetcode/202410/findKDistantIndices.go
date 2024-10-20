package main

func findKDistantIndices(nums []int, key int, k int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		s := 0
		e := 0
		if i-k < 0 {
			s = 0
		} else {
			s = i - k
		}
		if i+k > len(nums)-1 {
			e = len(nums) - 1
		} else {
			e = i + k
		}
		for s <= e {
			if abs(i-s) <= k && nums[s] == key {
				ans = append(ans, i)
				break
			}
			s++
		}
	}
	return ans
}
