package main

func minDeletionSize(strs []string) int {
	ans := 0
	for i := range strs[0] {
		for j := range strs {
			if j <= len(strs)-2 && strs[j][i] > strs[j+1][i] {
				ans++
				break
			}
		}
	}
	return ans
}
