package main

func checkRecord(s string) bool {
	ACount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			ACount++
		} else if s[i] == 'L' && i <= len(s)-3 {
			f := true
			for j := 1; j < 3; j++ {
				if s[i+j] != 'L' {
					f = false
				}
			}
			if f {
				return false
			}
		}
	}
	return ACount < 2
}

func main() {
	checkRecord("PPALLP")
}
