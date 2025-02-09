package main

func minOperations(logs []string) int {
	stack := []string{}
	for _, log := range logs {
		if log == "../" && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if log == "./" || (log == "../" && len(stack) == 0) {
			continue
		} else {
			stack = append(stack, log)
		}
	}
	return len(stack)
}
