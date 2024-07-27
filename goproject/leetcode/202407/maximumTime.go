package main

func maximumTime(time string) string {
	byteSlice := []byte(time)
	if byteSlice[0] == '?' {
		if time[1] >= '4' && time[1] <= '9' {
			byteSlice[0] = '1'
		} else {
			byteSlice[0] = '2'
		}
	}
	if byteSlice[1] == '?' {
		if byteSlice[0] == '2' {
			byteSlice[1] = '3'
		} else {
			byteSlice[1] = '9'
		}
	}
	if byteSlice[3] == '?' {
		byteSlice[3] = '5'
	}
	if byteSlice[4] == '?' {
		byteSlice[4] = '9'
	}
	return string(byteSlice)
}
