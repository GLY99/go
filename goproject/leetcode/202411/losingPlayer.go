package main

func losingPlayer(x int, y int) string {
	ops := min(x, y/4)
	if ops%2 == 0 {
		return "Bob"
	} else {
		return "Alice"
	}
}
