package main

func passThePillow(n int, time int) int {
	i := 1
	flag := true
	for time > 0 {
		time--
		if flag {
			i++
		} else {
			i--
		}
		if i == n || i == 1 {
			flag = !flag
		}
	}
	return i
}

func passThePillowI(n int, time int) int {
	// n = 4 time = 7 (n-1)*2一个循环，看最后走几个循环剩余多少时间
	// 0 1 2 3 4 5   6 7 8 9 10
	// 1 2 3 4 3 2   1 2 3 4 3
	time = time % ((n - 1) * 2)
	// 如果剩余的时间小于n，那么在剩余的时间过后，不会过队尾，但是会到队尾
	if time < n {
		return time + 1
	}
	// 否则就是过了队尾，那么剩余的时间需要减去前半段的耗时，然后用剩下的时间往回走
	return n - (time - (n - 1))
}
