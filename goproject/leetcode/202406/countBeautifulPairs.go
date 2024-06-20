package main

func countBeautifulPairs(nums []int) int {
	res := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		for nums[i] >= 10 {
			nums[i] = nums[i] / 10
		}
		for j := i + 1; j < length; j++ {
			if gcd(nums[i], nums[j]%10) == 1 {
				res++
			}
		}
	}
	return res
}

func gcd(x, y int) int {
	for y != 0 {
		tmp := x % y
		x = y
		y = tmp
	}
	return x
}
