package main

// https://leetcode.cn/problems/check-array-formation-through-concatenation/?envType=problem-list-v2&envId=hash-table

func canFormArray(arr []int, pieces [][]int) bool {
	mapping := make(map[int]int)
	for idx, piece := range pieces {
		mapping[piece[0]] = idx
	}
	for idx := 0; idx < len(arr); idx++ {
		num := arr[idx]
		if piece_idx, ok := mapping[num]; !ok {
			return false
		} else {
			piece_arr := pieces[piece_idx]
			for i := 0; i < len(piece_arr); i++ {
				if piece_arr[i] != arr[idx+i] {
					return false
				}
			}
			idx += len(piece_arr) - 1
		}
	}
	return true
}
