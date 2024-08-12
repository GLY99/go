package main

func removeOuterParentheses(s string) string {
	var stack []rune
	var ans []rune
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
			if len(stack) != 1 {
				ans = append(ans, c)
			}
		} else if c == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				ans = append(ans, c)
			}
		}
	}
	return string(ans)
}
