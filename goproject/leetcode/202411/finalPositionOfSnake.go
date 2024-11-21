package main

func finalPositionOfSnake(n int, commands []string) int {
	x, y := 0, 0
	commandsMapping := map[string][2]int{
		"UP": [2]int{-1, 0}, "DOWN": [2]int{1, 0},
		"LEFT": [2]int{0, -1}, "RIGHT": [2]int{0, 1},
	}
	for _, str := range commands {
		x += commandsMapping[str][0]
		y += commandsMapping[str][1]
	}
	return x*n + y
}
