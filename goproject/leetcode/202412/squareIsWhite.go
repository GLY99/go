package main

func squareIsWhite(coordinates string) bool {
	x := coordinates[0] - 'a' + 1
	y := coordinates[1] - '0'
	if (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0) {
		return false
	}
	return true
}
