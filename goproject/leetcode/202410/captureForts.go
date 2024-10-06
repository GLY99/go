package main

func captureForts(forts []int) int {
	ans := 0
	pre := -1
	for i := 0; i < len(forts); i++ {
		if forts[i] == 1 || forts[i] == -1 {
			if pre != -1 && forts[i] != forts[pre] {
				ans = max(ans, i-pre-1)
			}
			pre = i
		}
	}
	return ans
}
