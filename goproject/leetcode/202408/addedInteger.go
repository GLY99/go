package main

func addedInteger(nums1 []int, nums2 []int) int {
	sumNums1 := 0
	sumNums2 := 0
	length := len(nums1)
	for i := 0; i < length; i++ {
		sumNums1 += nums1[i]
		sumNums2 += nums2[i]
	}
	return (sumNums2 - sumNums1) / length
}
