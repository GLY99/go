package main

func maxDepth(s string) int {
	ans := 0
	count := 0
	var stack []rune = make([]rune, 0)
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
			count++
		}
		if c == ')' {
			stack = stack[:len(stack)-1]
			count--
		}
		ans = max(ans, count)
	}
	return ans
}
