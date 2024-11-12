package main

func isWinner(player1 []int, player2 []int) int {
	res1, res2 := score(player1), score(player2)
	if res1 > res2 {
		return 1
	} else if res1 < res2 {
		return 2
	} else {
		return 0
	}
}

func score(player []int) int {
	res := 0
	for i, x := range player {
		if i > 0 && player[i-1] == 10 || i > 1 && player[i-2] == 10 {
			res += 2 * x
		} else {
			res += x
		}
	}
	return res
}
