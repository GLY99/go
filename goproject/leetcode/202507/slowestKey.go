package main

// https://leetcode.cn/problems/slowest-key/?envType=problem-list-v2&envId=array

func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans := keysPressed[0]
	maxT := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		if releaseTimes[i]-releaseTimes[i-1] > maxT {
			ans = keysPressed[i]
			maxT = releaseTimes[i] - releaseTimes[i-1]
		} else if releaseTimes[i]-releaseTimes[i-1] == maxT && keysPressed[i] > ans {
			ans = keysPressed[i]
		}
	}
	return ans
}
