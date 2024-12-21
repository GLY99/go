package main

func minimumOperations(nums []int) int {
	mapping := make(map[int]interface{})
	for _, num := range nums {
		if num > 0 {
			mapping[num] = struct{}{}
		}
	}
	return len(mapping)
}
