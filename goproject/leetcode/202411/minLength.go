package main

func minLength(s string) int {
	var stack []byte = []byte{}
	for i := range s {
		stack = append(stack, s[i])
		stackLen := len(stack)
		if stackLen >= 2 && (string(stack[stackLen-2:]) == "AB" || string(stack[stackLen-2:]) == "CD") {
			stack = stack[:stackLen-2]
		}
	}
	return len(stack)
}
