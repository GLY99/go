package main

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	var ans int64 = 0
	s, e := 0, len(nums)-1
	for s <= e {
		if s < e {
			tmp, _ := strconv.Atoi(strconv.Itoa(nums[s]) + strconv.Itoa(nums[e]))
			ans += int64(tmp)
		} else {
			ans += int64(nums[e])
		}
		s++
		e--
	}
	return ans
}
