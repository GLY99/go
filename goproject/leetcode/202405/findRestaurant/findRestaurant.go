package main

import (
	"fmt"
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	res := make([]string, 0)
	mapping := make(map[string]int, 0)
	for idx, l1 := range list1 {
		mapping[l1] = idx
	}
	idxSum := math.MaxInt32
	for idx, l2 := range list2 {
		if val, ok := mapping[l2]; ok {
			if idx+val == idxSum {
				res = append(res, l2)
			} else if idx+val < idxSum {
				idxSum = idx + val
				res = []string{l2}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findRestaurant([]string{"a"}, []string{"b"}))
}
