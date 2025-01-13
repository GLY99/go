package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func fraction(cont []int) []int {
	x, y := 1, cont[len(cont)-1] // x分子 y分母
	for i := len(cont) - 2; i >= 0; i-- {
		newX := cont[i]*y + x
		newY := y

		// 化简
		g := gcd(newX, newY)
		x = newX / g
		y = newY / g

		x, y = y, x
	}
	x, y = y, x
	return []int{x, y}
}
