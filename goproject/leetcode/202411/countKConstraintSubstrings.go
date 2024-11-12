package main

func countKConstraintSubstrings(s string, k int) int {
	ans := 0
	n := len(s)
	cnt := [2]int{0, 0}
	left := 0
	for i := 0; i < n; i++ {
		cnt[s[i]&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[left]&1]--
			left++
		}
		ans += i - left + 1
	}
	return ans
}
