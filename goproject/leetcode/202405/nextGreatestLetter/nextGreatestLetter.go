package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	length := len(letters)
	left := 0
	right := length - 1
	if target >= letters[right] {
		return letters[0]
	}
	for left < right {
		mid := (left + right) / 2
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return letters[left]
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}
