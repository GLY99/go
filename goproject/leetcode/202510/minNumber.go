package main

// https://leetcode.cn/problems/form-smallest-number-from-two-digit-arrays/submissions/676136638/?envType=problem-list-v2&envId=array

func minNumber(nums1 []int, nums2 []int) int {
	mapping := make(map[int]int)
	nums1_min := nums1[0]
	nums2_min := nums2[0]
	for _, num := range nums1 {
		mapping[num]++
		nums1_min = min(num, nums1_min)
	}
	for _, num := range nums2 {
		mapping[num]++
		nums2_min = min(num, nums2_min)
	}
	ans := 10
	for k, v := range mapping {
		if v%2 == 0 {
			ans = min(ans, k)
		}
	}
	if ans != 10 {
		return ans
	}
	if nums1_min < nums2_min {
		return nums1_min*10 + nums2_min
	} else {
		return nums2_min*10 + nums1_min
	}
}
