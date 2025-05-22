package main

// https://leetcode.cn/problems/transform-array-by-parity/?envType=problem-list-v2&envId=sorting

func transformArray(nums []int) []int {
	a, b := 0, 0
	for _, num := range nums {
		if num%2 == 0 {
			a++
		} else {
			b++
		}
	}
	ans := make([]int, 0)
	for a != 0 || b != 0 {
		if a != 0 {
			ans = append(ans, 0)
			a--
			continue
		}
		if b != 0 {
			ans = append(ans, 1)
			b--
			continue
		}
	}
	return ans
}
