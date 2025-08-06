package main

// https://leetcode.cn/problems/two-out-of-three/?envType=problem-list-v2&envId=array

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	mapping := make(map[int]int)
	for idx, nums := range [][]int{nums1, nums2, nums3} {
		for _, num := range nums {
			mapping[num] |= 1 << idx
		}
	}
	// v & (v - 1)解释
	// 1.三个数组出现：111&110=110(√)
	// 2.两个数组出现：110&101=100(√),101&100=100(√),011&010=010(√)
	// 3.一个数组出现：100&011=000(×),010&001=000(×),001&000=000(×),
	ans := make([]int, 0)
	for num, idx := range mapping {
		if idx&(idx-1) > 0 {
			ans = append(ans, num)
		}
	}
	return ans
}
