package main

// https://leetcode.cn/problems/prime-in-diagonal/?envType=daily-question&envId=2025-03-18

func diagonalPrime(nums [][]int) int {
	length := len(nums)
	ans := 0
	for i := 0; i < length; i++ {
		if isPrime(nums[i][i]) {
			ans = max(ans, nums[i][i])
		}
		if isPrime(nums[i][length-i-1]) {
			ans = max(ans, nums[i][length-i-1])
		}
	}
	return ans
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
