package main

func diStringMatch(s string) []int {
	length := len(s)
	ans := make([]int, length+1)
	left := 0
	right := length
	for idx, v := range s {
		if v == 'I' {
			ans[idx] = left
			left++
		} else if v == 'D' {
			ans[idx] = right
			right--
		}
	}
	ans[length] = right
	return ans
}
