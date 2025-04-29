package main

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	if prices[0]+prices[1] > money {
		return money
	}
	return money - (prices[0] + prices[1])
}
