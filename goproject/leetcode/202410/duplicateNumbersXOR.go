package main

func duplicateNumbersXOR(nums []int) int {
	var mapping map[int]interface{} = make(map[int]interface{})
	res := 0
	for _, num := range nums {
		if _, ok := mapping[num]; ok {
			res ^= num
		} else {
			mapping[num] = struct{}{}
		}
	}
	return res
}
