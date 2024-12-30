package main

func minimumChairs(s string) int {
	var ans int
	var num int
	for _, c := range s {
		if c == 'E' {
			num++
		} else {
			num--
		}
		ans = max(num, ans)
	}
	return ans
}
