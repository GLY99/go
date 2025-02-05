package main

func makeGood(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if s[i]-'a' == stack[len(stack)-1]-'A' || s[i]-'A' == stack[len(stack)-1]-'a' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	ans := string(stack)
	return ans
}
