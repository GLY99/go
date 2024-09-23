package main

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	// 用mapping记录每个人信任的人数以及被人信任的人数
	var mapping map[int]*[2]int = map[int]*[2]int{}
	for _, t := range trust {
		idx := t[0]
		trustIdx := t[1]
		if _, ok := mapping[idx]; !ok {
			mapping[idx] = &[2]int{0, 0}
		}
		if _, ok := mapping[trustIdx]; !ok {
			mapping[trustIdx] = &[2]int{0, 0}
		}
		(*mapping[idx])[0]++
		(*mapping[trustIdx])[1]++
	}
	for k, v := range mapping {
		if v[0] == 0 && v[1] == n-1 {
			return k
		}
	}
	return -1
}
