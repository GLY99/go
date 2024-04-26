package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
**/
var pick = 6

func guess(num int) int {
	if num == pick {
		return 0
	} else if num > pick {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) >> 1
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if guess(l) == 0 {
		return l
	} else {
		return 0
	}
}

func main() {
	guessNumber(10)
}
