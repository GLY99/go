package main

// https://leetcode.cn/problems/path-crossing/?envType=problem-list-v2&envId=hash-table

func isPathCrossing(path string) bool {
	mapping := make(map[[2]int]interface{})
	x, y := 0, 0
	mapping[[2]int{x, y}] = struct{}{}
	for i := 0; i < len(path); i++ {
		if path[i] == 'N' {
			y++
		} else if path[i] == 'S' {
			y--
		} else if path[i] == 'E' {
			x++
		} else {
			x--
		}
		if _, ok := mapping[[2]int{x, y}]; ok == true {
			return true
		} else {
			mapping[[2]int{x, y}] = struct{}{}
		}
	}
	return false
}
