package main

func finalValueAfterOperations(operations []string) int {
	var res int = 0
	for _, opoperation := range operations {
		if opoperation == "X++" || opoperation == "++X" {
			res++
		} else {
			res--
		}
	}
	return res
}
