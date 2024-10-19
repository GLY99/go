package main

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	ansNums := make([][]int, 0)
	a := 0
	b := 0
	for a <= len(nums1) || b <= len(nums2) {
		if a == len(nums1) {
			ansNums = append(ansNums, nums2[b:]...)
			break
		}
		if b == len(nums2) {
			ansNums = append(ansNums, nums1[a:]...)
			break
		}
		if nums1[a][0] < nums2[b][0] {
			ansNums = append(ansNums, nums1[a])
			a++
		} else if nums1[a][0] > nums2[b][0] {
			ansNums = append(ansNums, nums2[b])
			b++
		} else {
			ansNums = append(ansNums, []int{nums1[a][0], nums1[a][1] + nums2[b][1]})
			a++
			b++
		}
	}
	return ansNums
}
