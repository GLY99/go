package main

// https://leetcode.cn/problems/find-the-distinct-difference-array/?envType=problem-list-v2&envId=array

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	mapping := map[int]struct{}{}
	counter := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		mapping[nums[i]] = struct{}{}
		counter[i] = len(mapping)
	}
	mapping = map[int]struct{}{}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		mapping[nums[i]] = struct{}{}
		res[i] = len(mapping) - counter[i+1]
	}
	return res
}
