package main

import "fmt"

func quickSort(start, end int, nums []int) {
	left := start
	right := end
	flag := nums[(left+right)/2]
	for left < right {
		// 从左边找到大于等于flag的值
		for left <= end && nums[left] < flag {
			left++
		}
		// 从右边找到小于等于flag的值
		for right >= start && nums[right] > flag {
			right--
		}
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	if left == right {
		left++
		right--
	}

	if left < end {
		quickSort(left, end, nums)
	}
	if right > start {
		quickSort(start, right, nums)
	}
}

func main() {
	nums := []int{-9, 78, 0, 23, -567, 70}
	quickSort(0, len(nums)-1, nums)
	fmt.Println(nums)
}
