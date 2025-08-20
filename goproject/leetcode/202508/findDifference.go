package main

// https://leetcode.cn/problems/find-the-difference-of-two-arrays/?envType=problem-list-v2&envId=array

func findDifference(nums1 []int, nums2 []int) [][]int {
	counter1 := make(map[int]interface{})
	counter2 := make(map[int]interface{})
	for _, num := range nums1 {
		counter1[num] = struct{}{}
	}
	for _, num := range nums2 {
		counter2[num] = struct{}{}
	}
	ans := make([][]int, 2)
	for num, _ := range counter1 {
		if _, ok := counter2[num]; !ok {
			ans[0] = append(ans[0], num)
		}
	}
	for num, _ := range counter2 {
		if _, ok := counter1[num]; !ok {
			ans[1] = append(ans[1], num)
		}
	}
	return ans
}
