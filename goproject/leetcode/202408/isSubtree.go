package main

import (
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, subRoot) || check(root.Left, subRoot) || check(root.Right, subRoot)
}

func check(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return check(root1.Left, root2.Left) && check(root1.Right, root2.Right)
	}
	return false
}

// 先序遍历+字符串匹配
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	maxEle := math.MinInt32
	getMaxElement(root, &maxEle)
	getMaxElement(subRoot, &maxEle)
	lNull := maxEle + 1
	rNull := maxEle + 2
	rl, sl := getDfsOrder(root, []int{}, lNull, rNull), getDfsOrder(subRoot, []int{}, lNull, rNull)
	return isSub(rl, sl)
}

func getMaxElement(root *TreeNode, maxEle *int) {
	if root == nil {
		return
	}
	if root.Val > *maxEle {
		*maxEle = root.Val
	}
	getMaxElement(root.Left, maxEle)
	getMaxElement(root.Right, maxEle)
}

func getDfsOrder(root *TreeNode, list []int, lNull, rNull int) []int {
	if root == nil {
		return list
	}
	list = append(list, root.Val)
	if root.Left != nil {
		list = getDfsOrder(root.Left, list, lNull, rNull)
	} else {
		list = append(list, lNull)
	}

	if root.Right != nil {
		list = getDfsOrder(root.Right, list, lNull, rNull)
	} else {
		list = append(list, rNull)
	}
	return list
}

func isSub(s, t []int) bool {
	sLen, tLen := len(s), len(t)
	for i := 0; i < sLen; i++ {
		if s[i] == t[0] {
			flag := true
			for j := 0; j < tLen; j++ {
				if s[i+j] != t[j] {
					flag = false
					break
				}
			}
			if flag {
				return flag
			}
		}
	}
	return false
}
