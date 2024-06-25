package main

func tictactoe(moves [][]int) string {
	mapping := new([3][3]int)
	flag := 1
	for _, points := range moves {
		mapping[points[0]][points[1]] = flag
		flag = -flag
	}
	if abs(mapping[0][0], mapping[1][1], mapping[2][2]) == 3 || abs(mapping[0][2], mapping[1][1], mapping[2][0]) == 3 {
		if mapping[1][1] == 1 {
			return "A"
		}
		return "B"
	}
	for i := 0; i < 3; i++ {
		if abs(mapping[i][0], mapping[i][1], mapping[i][2]) == 3 || abs(mapping[0][i], mapping[1][i], mapping[2][i]) == 3 {
			if mapping[i][i] == 1 {
				return "A"
			}
			return "B"
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func abs(args ...int) int {
	ans := 0
	for _, arg := range args {
		ans += arg
	}
	if ans < 0 {
		return -ans
	}
	return ans
}
