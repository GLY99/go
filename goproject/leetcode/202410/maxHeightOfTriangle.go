package main

func maxHeightOfTriangle(red int, blue int) int {
	return max(maxHeight(red, blue), maxHeight(blue, red))
}

func maxHeight(x, y int) int {
	for i := 1; ; i++ {
		if i%2 == 0 {
			x -= i
			if x < 0 {
				return i - 1
			}
		} else {
			y -= i
			if y < 0 {
				return i - 1
			}
		}
	}
}
