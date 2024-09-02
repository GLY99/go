package main

func reverseStr(s string, k int) string {
	length := len(s)
	count := length / (2 * k)
	sSlice := []byte(s)
	for i := 0; i < count; i++ {
		start := i * 2 * k
		end := start + k - 1
		for start < end {
			sSlice[start], sSlice[end] = sSlice[end], sSlice[start]
			start++
			end--
		}
	}
	if length%(2*k) < k {
		start := count * 2 * k
		end := length - 1
		for start < end {
			sSlice[start], sSlice[end] = sSlice[end], sSlice[start]
			start++
			end--
		}
	} else {
		start := count * 2 * k
		end := start + k - 1
		for start < end {
			sSlice[start], sSlice[end] = sSlice[end], sSlice[start]
			start++
			end--
		}
	}
	return string(sSlice)
}
