package main

func duplicateZeros(arr []int) {
	count := 0
	idx := 0
	length := len(arr)
	for count < length {
		if arr[idx] == 0 {
			count += 2
		} else {
			count++
		}
		if count >= length {
			break
		}
		idx++
	}
	j := length - 1
	if count == length+1 {
		arr[j] = 0
		j--
		idx--
	}
	for j >= 0 {
		arr[j] = arr[idx]
		j--
		if j < 0 {
			break
		}
		if arr[idx] == 0 {
			arr[j] = 0
			j--
		}
		idx--
	}
}
