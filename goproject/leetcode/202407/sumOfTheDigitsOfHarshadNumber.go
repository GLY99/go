package main

func sumOfTheDigitsOfHarshadNumber(x int) int {
	res := 0
	tmpX := x
	for tmpX > 0 {
		res += tmpX % 10
		tmpX = tmpX / 10
	}
	if x%res == 0 {
		return res
	}
	return -1
}
