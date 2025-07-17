package main

// https://leetcode.cn/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	return quickSelect(nums, 0, length-1, length-k)
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	pos := partition(nums, l, r)
	if k == pos {
		return nums[pos]
	} else if k < pos {
		return quickSelect(nums, l, pos-1, k)
	} else {
		return quickSelect(nums, pos+1, r, k)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[l]
	i := l
	j := r
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		if l >= r {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[l] = nums[l], nums[i]
	return i
}
