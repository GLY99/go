package main

func maxDivScore(nums []int, divisors []int) int {
	res := 0
	count := -1
	for _, divisor := range divisors {
		tmp := 0
		for _, num := range nums {
			if num%divisor == 0 {
				tmp++
			}
		}
		if tmp > count || (tmp == count && divisor < res) {
			count = tmp
			res = divisor
		}
	}
	return res
}

func main() {
	var nums = []int{1, 2, 3}
	var divisors = []int{2, 3, 4}
	maxDivScore(nums, divisors)
}
