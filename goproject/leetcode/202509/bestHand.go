package main

import "bytes"

// https://leetcode.cn/problems/best-poker-hand/?envType=problem-list-v2&envId=array

func bestHand(ranks []int, suits []byte) string {
	if bytes.Count(suits, suits[:1]) == 5 {
		return "Flush"
	}
	cnt, pair := map[int]int{}, false
	for _, rank := range ranks {
		cnt[rank]++
		if cnt[rank] == 3 {
			return "Three of a Kind"
		}
		if cnt[rank] == 2 {
			pair = true
		}
	}
	if pair {
		return "Pair"
	}
	return "High Card"
}
