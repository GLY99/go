package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return Helper(nums, 0, len(nums)-1)
}

func Helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = Helper(nums, left, mid-1)
	root.Right = Helper(nums, mid+1, right)
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	node := sortedArrayToBST(nums)
	fmt.Println(node)
}
