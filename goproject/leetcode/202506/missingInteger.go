package main

// https://leetcode.cn/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/?envType=problem-list-v2&envId=sorting

func missingInteger(nums []int) int {
	mapping := make(map[int]interface{})
	for _, num := range nums {
		mapping[num] = struct{}{}
	}
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}
	_, ok := mapping[sum]
	for ok {
		sum++
		_, ok = mapping[sum]
	}
	return sum
}
