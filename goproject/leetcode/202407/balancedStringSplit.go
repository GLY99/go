package main

func balancedStringSplit(s string) int {
	ans := 0
	d := 0
	for _, ch := range s {
		if ch == 'L' {
			d++
		} else {
			d--
		}
		if d == 0 {
			ans++
		}
	}
	return ans
}
