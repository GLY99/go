package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/maximum-difference-by-remapping-a-digit/?envType=daily-question&envId=2025-06-14

func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	t := s
	idx := 0
	for idx < len(s) && s[idx] == '9' {
		idx++
	}
	if idx < len(s) {
		s = strings.ReplaceAll(s, string(s[idx]), "9")
	}
	t = strings.ReplaceAll(t, string(t[0]), "0")
	max, _ := strconv.Atoi(s)
	min, _ := strconv.Atoi(t)
	return max - min
}
