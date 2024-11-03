package main

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(index))
	for i := 0; i < len(index); i++ {
		insert(target, i, index[i], nums[i])
	}
	return target
}

func insert(target []int, curLength int, idx int, num int) {
	for i := curLength - 1; i >= idx; i-- {
		target[i+1] = target[i]
	}
	target[idx] = num
}
