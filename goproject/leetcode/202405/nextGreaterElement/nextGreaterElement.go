package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	num1Length := len(nums1)
	num2Length := len(nums2)
	res := make([]int, num1Length)
	for i, num1 := range nums1 {
		find := false
		flag := false
		for j := 0; j < num2Length; j++ {
			if nums2[j] == num1 {
				flag = true
			}
			if flag && nums2[j] > num1 {
				res[i] = nums2[j]
				find = true
				break
			}
		}
		if !find {
			res[i] = -1
		}
	}
	return res
}

func main() {
	nextGreaterElement([]int{1, 2}, []int{2, 3})
}
