package main

func mostVisited(n int, rounds []int) []int {
	start, end := rounds[0], rounds[len(rounds)-1]
	ret := make([]int, 0)
	if start <= end {
		// 跑了x圈，最后一圈从start - end
		for i := start; i <= end; i++ {
			ret = append(ret, i)
		}
	} else {
		for i := 1; i <= end; i++ {
			ret = append(ret, i)
		}
		for i := start; i <= n; i++ {
			ret = append(ret, i)
		}
	}
	return ret
}
