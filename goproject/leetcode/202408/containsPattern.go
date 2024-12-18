package main

func containsPattern(arr []int, m int, k int) bool {
	// arr = [2,2,1,2,2,1,1,1,2,1], m = 2, k = 2
	length := len(arr)
	for i := 0; i < length; i++ {
		count := 1
		for j := i + m; j <= length-m; j += m {
			isSame := true
			for t := 0; t < m; t++ {
				if arr[j+t] != arr[i+t] {
					isSame = false
					break
				}
			}
			if isSame {
				count++
				if count >= k {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}

func containsPatternI(arr []int, m int, k int) bool {
	length := len(arr)
	for i := 0; i <= length-m*k; i++ {
		offset := 0
		for offset < m*k {
			if arr[i+offset] != arr[i+offset%m] {
				break
			}
			offset++
		}
		if offset == m*k {
			return true
		}
	}
	return false
}
