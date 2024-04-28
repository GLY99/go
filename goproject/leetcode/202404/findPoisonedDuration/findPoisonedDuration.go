package main

// https://leetcode.cn/problems/teemo-attacking/

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	expired := 0
	for _, time := range timeSeries {
		if time >= expired {
			ans += duration
		} else {
			ans += time + duration - expired
		}
		expired = time + duration
	}
	return ans
}

func main() {
	findPoisonedDuration([]int{1, 2}, 2)
}
