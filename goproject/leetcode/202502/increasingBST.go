package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			vals = append(vals, node.Val)
			inorder(node.Right)
		}
	}
	inorder(root)
	newRoot := &TreeNode{}
	curNode := newRoot
	for _, val := range vals {
		curNode.Right = &TreeNode{Val: val}
		curNode = curNode.Right
	}
	return newRoot.Right
}

func increasingBST1(root *TreeNode) *TreeNode {
	newRoot := &TreeNode{}
	curNode := newRoot
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		curNode.Right = node
		node.Left = nil
		curNode = node
		inorder(node.Right)
	}
	inorder(root)
	return newRoot.Right
}
