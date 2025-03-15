package main

func f1(a byte, b byte) int {
	if a > b {
		return int(a - b)
	} else {
		return int(b - a)
	}
}

func scoreOfString(s string) int {
	var ans int
	for i := 1; i < len(s); i++ {
		ans += f1(s[i], s[i-1])

	}
	return ans
}
