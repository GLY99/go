package main

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	comb := []int{}
	var dfs func(int, int)
	dfs = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, comb...))
			return
		}
		// 不选当前的数字
		dfs(target, idx+1)
		// 选择当前的数字
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	combinationSum(candidates, target)
}
