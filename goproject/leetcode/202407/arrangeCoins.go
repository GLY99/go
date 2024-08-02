package main

func arrangeCoins(n int) int {
	left, right := 0, n
	for left <= right {
		// 逼近右边界 l=1 r=2 m取2而不是1
		mid := left + (right-left+1)/2
		if mid*(mid+1) == 2*n {
			return mid
		} else if mid*(mid+1) < 2*n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}
