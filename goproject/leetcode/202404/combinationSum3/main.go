package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	tmp := []int{}
	var dfs func(int, int)
	dfs = func(cur int, target int) {
		if target == 0 {
			if len(tmp) == k {
				res = append(res, append([]int{}, tmp...))
				return
			}
		}
		if cur > 9 || cur > target || len(tmp) > k {
			return
		}
		tmp = append(tmp, cur)
		dfs(cur+1, target-cur)
		tmp = tmp[:len(tmp)-1]
		dfs(cur+1, target)
	}
	dfs(1, n)
	return res
}

func main() {
	// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
	// 只使用数字1到9
	// 每个数字 最多使用一次
	// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
	// 2 <= k <= 9
	// 1 <= n <= 60
	k := 3
	n := 9
	combs := combinationSum3(k, n)
	fmt.Println(combs)
}
