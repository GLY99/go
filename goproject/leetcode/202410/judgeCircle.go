package main

func judgeCircle(moves string) bool {
	var x int = 0
	var y int = 0
	for _, c := range moves {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}
	return x == 0 && y == 0
}
