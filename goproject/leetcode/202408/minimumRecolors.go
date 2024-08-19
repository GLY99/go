package main

func minimumRecolors(blocks string, k int) int {
	length := len(blocks)
	l := 0
	r := 0
	BCount := 0
	ans := 0
	for r < length {
		if r >= k {
			if blocks[l] == 'B' {
				BCount--
			}
			l++
		}
		if blocks[r] == 'B' {
			BCount++
		}
		ans = max(ans, BCount)
		r++
	}
	return k - ans
}
