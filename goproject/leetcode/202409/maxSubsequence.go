package main

import "sort"

func maxSubsequence(nums []int, k int) []int {
	ids := make([]int, len(nums))
	for i := range ids {
		ids[i] = i
	}
	// 将源数据的下标按照对应的值从大到小排序
	sort.Slice(ids, func(i, j int) bool { return nums[ids[i]] > nums[ids[j]] })
	// 排序后前k个下标对应的值就是最大的，将下标按照顺序再次排序
	sort.Ints(ids[:k])
	// 记录答案
	ans := make([]int, k)
	for i, j := range ids[:k] {
		ans[i] = nums[j]
	}
	return ans
}
