package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	graph := map[int][]int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		for _, child := range []*TreeNode{node.Left, node.Right} {
			if child != nil {
				graph[node.Val] = append(graph[node.Val], child.Val)
				graph[child.Val] = append(graph[child.Val], node.Val)
				dfs(child)
			}
		}
	}
	dfs(root)
	q := [][]int{{start, 0}}
	visited := map[int]bool{}
	visited[start] = true
	time := 0
	for len(q) > 0 {
		arr := q[0]
		q = q[1:]
		nodeVal := arr[0]
		time = arr[1]
		for _, childVal := range graph[nodeVal] {
			if !visited[childVal] {
				q = append(q, []int{childVal, time + 1})
				visited[childVal] = true
			}
		}
	}
	return time
}
