package main

func numberOfChild(n int, k int) int {
	// (n - 1) * 2为一个循环
	x := k % ((n - 1) * 2)
	if x <= n-1 {
		return x
	} else {
		return 2*(n-1) - x
	}
}
