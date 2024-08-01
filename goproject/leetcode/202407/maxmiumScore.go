package main

import "sort"

func maxmiumScore(cards []int, cnt int) int {
	ans := 0
	tmp := 0
	odd, even := -1, -1
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	for i := 0; i < cnt; i++ {
		tmp += cards[i]
		if cards[i]%2 == 1 {
			odd = cards[i]
		} else {
			even = cards[i]
		}
	}
	if tmp%2 == 0 {
		return tmp
	}
	for i := cnt; i < len(cards); i++ {
		if cards[i]%2 == 1 {
			if even != -1 {
				ans = max(ans, tmp-even+cards[i])
			}
		} else {
			if odd != -1 {
				ans = max(ans, tmp-odd+cards[i])
			}
		}
	}
	return ans
}
