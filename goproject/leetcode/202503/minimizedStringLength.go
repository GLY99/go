package main

func minimizedStringLength(s string) int {
	var mask uint
	for _, c := range s {
		mask |= 1 << (c - 'a')
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if mask&(1<<i) != 0 {
			ans++
		}
	}
	return ans
}
