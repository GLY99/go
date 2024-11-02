package main

import "math/bits"

func minChanges(n, k int) (res int) {
	for n > 0 || k > 0 {
		if n&1 == 0 && k&1 == 1 {
			return -1
		}
		if n&1 == 1 && k&1 == 0 {
			res++
		}
		k >>= 1
		n >>= 1
	}
	return
}

func minChanges1(n, k int) (res int) {
	// n:1 k:0 t
	// n:1 k:1 t
	// n:0 k:1 f
	// n:0 k:0 t
	if n&k != k {
		res = -1
	} else {
		res = bits.OnesCount(uint(n ^ k))
	}
	return
}
