package main

func numberOfMatches(n int) int {
	var res int = 0
	for n > 1 {
		res += n / 2
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n/2 + 1
		}
	}
	return res
}
