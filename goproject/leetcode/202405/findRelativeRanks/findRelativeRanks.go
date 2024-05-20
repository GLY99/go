package main

import (
	"sort"
	"strconv"
)

var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
	mapping := map[int]int{}
	for idx, s := range score {
		mapping[s] = idx
	}
	res := make([]string, len(score))
	sort.Slice(score, func(i, j int) bool { return score[i] > score[j] })
	for idx1, s := range score {
		idx2 := mapping[s]
		if idx1 < 3 {
			res[idx2] = desc[idx1]
		} else {
			res[idx2] = strconv.Itoa(idx1 + 1)
		}
	}
	return res
}

func main() {
	findRelativeRanks([]int{1, 2, 3})
}
