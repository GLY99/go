package main

import "math"

func canAliceWin(n int) bool {
	if n < 10 {
		return false
	}
	x := 10
	n -= x
	for n >= x-1 {
		n -= x - 1
		x--
	}
	if (x-1)%2 == 0 {
		return false
	}
	return true
}

func canAliceWinI(n int) bool {
	// 10+9+8+⋯+(11−x)=(21−x)x/2<=n
	// x <= (21 - sqrt(21^2 - 8n))/2
	x := (21 - int(math.Ceil(math.Sqrt(float64(441-n*8))))) / 2
	return x%2 != 0
}
