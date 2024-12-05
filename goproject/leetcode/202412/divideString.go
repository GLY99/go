package main

func divideString(s string, k int, fill byte) []string {
	strs := make([]string, 0)
	x := len(s) % k
	if x != 0 {
		for i := 0; i < k-x; i++ {
			s += string(fill)
		}
	}
	for i := 0; i < len(s); i += k {
		strs = append(strs, s[i:i+k])
	}
	return strs
}
