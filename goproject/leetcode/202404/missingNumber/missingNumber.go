package main

// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数
// nums 中的所有数字都 独一无二
// n == nums.length

func missingNumber(nums []int) int {
	length := len(nums)
	total := length * (length + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}

func main() {
	missingNumber([]int{0, 1})
}
