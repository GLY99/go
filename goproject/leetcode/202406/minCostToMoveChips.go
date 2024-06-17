package main

func minCostToMoveChips(position []int) int {
	ans := new([2]int)
	for _, p := range position {
		ans[p%2]++
	}
	return min(ans[0], ans[1])
}
