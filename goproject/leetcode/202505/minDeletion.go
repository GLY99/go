package main

import "slices"

// https://leetcode.cn/problems/minimum-deletions-for-at-most-k-distinct-characters/?envType=problem-list-v2&envId=sorting

func minDeletion(s string, k int) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	slices.Sort(cnt[:])
	ans := 0
	for _, c := range cnt[:26-k] {
		ans += c
	}
	return ans
}
