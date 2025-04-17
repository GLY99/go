package main

func countPairs(nums []int, k int) int {
	mapping := make(map[int][]int)
	ans := 0
	for idx, num := range nums {
		if idxs, ok := mapping[num]; ok {
			for _, i := range idxs {
				if (i*idx)%k == 0 {
					ans++
				}
			}
		}
		mapping[num] = append(mapping[num], idx)
	}
	return ans
}
