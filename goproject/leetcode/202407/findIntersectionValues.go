package main

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	ans := make([]int, 2)
	mapping1 := make(map[int]int)
	mapping2 := make(map[int]int)
	for _, num := range nums1 {
		mapping1[num] += 1
	}
	for _, num := range nums2 {
		mapping2[num] += 1
		_, ok := mapping1[num]
		if ok {
			ans[1] += 1
		}
	}
	for _, num := range nums1 {
		if _, ok := mapping2[num]; ok {
			ans[0] += 1
		}
	}
	return ans
}
