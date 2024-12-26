package main

func isSubstringPresent(s string) bool {
	mapping := make(map[string]interface{})
	byteSlice := []byte(s)
	for i := 0; i < len(byteSlice)-1; i++ {
		if byteSlice[i] == byteSlice[i+1] {
			return true
		}
		tmpStr := string(byteSlice[i : i+2])
		tmpStrReverse := string(byteSlice[i+1]) + string(byteSlice[i])
		if _, ok := mapping[tmpStrReverse]; ok {
			return true
		} else {
			mapping[tmpStr] = struct{}{}
		}
	}
	return false
}
