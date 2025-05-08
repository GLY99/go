package main

func isUnique(astr string) bool {
	mapping := make(map[byte]interface{})
	for i := 0; i < len(astr); i++ {
		if _, ok := mapping[astr[i]]; ok {
			return false
		}
		mapping[astr[i]] = struct{}{}
	}
	return true
}
